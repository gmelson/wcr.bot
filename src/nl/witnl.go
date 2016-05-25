// witnl
package nl

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const(
	var REPLACE_WCR = "Weekend Coffee Roasters"
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


func (m *Message) ReplaceWords(q string) (err error){
	
}
