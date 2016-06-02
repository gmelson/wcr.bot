// wcr.bot project main.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"nl"
	"appengine"
)

type Body struct {
	Object string `json:"object"`
	Entry  []struct {
		ID        string `json:"id"`
		Time      int64  `json:"time"`
		Messaging []struct {
			Sender struct {
				ID string `json:"id"`
			} `json:"sender"`
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Timestamp int64 `json:"timestamp"`
			Message   struct {
				Mid  string `json:"mid"`
				Seq  int    `json:"seq"`
				Text string `json:"text"`
			} `json:"message"`
		} `json:"messaging"`
	} `json:"entry"`
}

func init() {
	http.HandleFunc("/webhook", handler)
}

func handler(w http.ResponseWriter, req *http.Request) {

	c := appengine.NewContext(req)
	
	var answer string
	var body []byte
	var err error

	//Handle messenger challenge if GET otherwise handle message
	if req.Method == "GET" {
		if req.URL.Query().Get("hub.verify_token") == "i_like_a_good_challenge" {
			answer = req.URL.Query().Get("hub.challenge")
		} else {
			answer = "Error, wrong validation token"
		}
	} else {
		c.Infof("Handling POST method")

		//This is a POST so expect FB Message
		body, err = ioutil.ReadAll(req.Body)
		if nil == err {
			c.Infof("No error trying to parse body")
			bodyParts := new(Body)
			err = json.Unmarshal(body, bodyParts)
			if nil == err {
				c.Infof("No error unmarshalling bodyparts")

				for _, messages := range bodyParts.Entry[0].Messaging {
					
					c.Infof("Getting message")

					// Substitute words and get Intent with Entities
					message := new(nl.Message)
					message.Context = c
					_, messages.Message.Text = message.ReplaceWords(messages.Message.Text)
					err = message.GetIntent(messages.Message.Text)
					if err == nil {
						c.Infof("Getting answer")

						//Answer based on rules
						a := new(nl.Answer)
						a.Message = *message
						err, answer = a.GetAnswer()
						c.Infof("Get ready to send answer " + answer)
						
						//send a resposne
						err = message.SendResponse(answer,messages.Sender.ID)
					}

				}

			}

		}

	}

	if nil != err {
			c.Infof("Error:" + err.Error())
	}
	
	//Set for cross domain stuff
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Write([]byte(answer))
}

