// witnl
package nl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"appengine/urlfetch"
	"appengine"
	"bytes"
)

const (
	WCR_REPLACE_1 = "weekend coffee roasters"
	WCR_REPLACE_2 = "weekend coffee"
	token = "EAAQFA0PF8WsBABX4RJWOlP5KKWIgS6tJMZA20LfZC9B35EAZCnNyIfpCVWQCSDDlS2ukb33rd3ns3ZA0TZAt09gSIu7qbIpeyjNhAy1y4Xs8XnIjeZCaZB5FH84IAdjeNZB3aZBxugGg9muabtksdjonVZCcfaZARfgNIJXZAwQlx1XxOwZDZD"
	postUrl = "https://graph.facebook.com/v2.6/me/messages"

)

type Message struct {
	_text    string `json:"_text"`
	Entities struct {
		Intent []struct {
			Confidence float64 `json:"confidence"`
			Value      string  `json:"value"`
		} `json:"intent"`
		PhraseToTranslate []struct {
			Confidence float64 `json:"confidence"`
			Suggested  bool    `json:"suggested"`
			Type       string  `json:"type"`
			Value      string  `json:"value"`
		} `json:"phrase_to_translate"`
	} `json:"entities"`
	MsgID string `json:"msg_id"`
	Context appengine.Context
}

func (m *Message) GetIntent(q string) (err error) {

	// setup up query and URL parameters
	// encode before creating a new request
	requestUrl, _ := url.Parse("https://api.wit.ai/message")
	params := url.Values{}
	params.Add("v", "20160523")
	params.Add("q", q)
	requestUrl.RawQuery = params.Encode()
	req, err := http.NewRequest("GET", requestUrl.String(), nil)
	if err != nil {
		fmt.Printf("GET ERROR=%s\n", err)
		return
	}

	// Set Bearer provided by Wit.ai and make GET request
	req.Header.Add("Authorization", " Bearer AN7H6G4KTGHA7LFKEJ6AAZZRODLTZTFR")
	//client := &http.Client{}
	client := urlfetch.Client(m.Context)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request")
		return
	}

	// Read body and parse json into object
	jsonResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("READ ERROR=%s\n", err)
		return
	}
	fmt.Printf("json returned = %s\n", jsonResponse)
	err = json.Unmarshal(jsonResponse, m)
	if err != nil {
		fmt.Printf("JSON error %s\n", err)
		return
	}

	return

}

func (m *Message) ReplaceWords(q string) (err error, replaced string) {

	//Order dependant; Replace full Weekend Coffee Roasters, then look for Weekend Coffee
	// followed by Weekend Roasters
	q = strings.ToLower(q)
	if strings.Contains(q, WCR_REPLACE_1) {
		replaced = strings.Replace(q, WCR_REPLACE_1, "wcr", -1)
	} else if strings.Contains(q, WCR_REPLACE_2) {
		replaced = strings.Replace(q, WCR_REPLACE_2, "wcr", -1)
	} else {
		replaced = q
	}

	return

}

func (m *Message) SendResponse(answer string, sender string) (err error) {
	
	requestUrl, _ := url.Parse(postUrl)
	params := url.Values{}
	params.Add("access_token", token)
	requestUrl.RawQuery = params.Encode()

	//create json response
	messageData := fmt.Sprintf("{\"text\":\"%s\"}", answer)
	m.Context.Infof("messageData: " + messageData)
	str := fmt.Sprintf("{\"recipient\":{\"id\":\"%s\"},\"message\":%s}", sender, messageData)
	m.Context.Infof("str: " + str)
	var jsonStr = []byte(str)	
	req, err := http.NewRequest("POST", requestUrl.String(), bytes.NewBuffer(jsonStr))
	if err != nil {
		fmt.Printf("POST ERROR=%s\n", err)
		return
	}

	// encode before creating a new request
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Add("qs", fmt.Sprintf("{\"%s\"}", token) )

	client := urlfetch.Client(m.Context)
	_, err = client.Do(req)
	if err != nil {
		fmt.Printf("Error making request")
		return
	}
	return
}