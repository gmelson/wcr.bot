// witnl
package nl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	WCR_REPLACE_1 = "weekend coffee roasters"
	WCR_REPLACE_2 = "weekend coffee"
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
	client := &http.Client{}
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
