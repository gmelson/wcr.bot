// wcr.bot project main.go
package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"nl"
	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
	"strings"
	"time"
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
	http.HandleFunc("/tagging", taghandler)
}

/** 
	WCR tag handler
	Look for existing data object; if found, check timestamp to see if it is
	within the threshold specified (also contained in data object).
	Otherwise, set the timestamp and continue.
	Set for cross domain stuff
*/
func taghandler(w http.ResponseWriter, req *http.Request) {
	c := appengine.NewContext(req)
	if req.Method == "GET" && strings.Contains(req.URL.Path, "tagging") {
		visitorId := req.URL.Query().Get("visitor.id")
		pageId := req.URL.Query().Get("page.id")
		log.Infof(c,"The visitor id= " + visitorId)
	log.Infof(c,"The page id= " + pageId)

		var err error
		var offer = "false"
		
		tag := &nl.Tag{
        	Id: visitorId,
        	LastVisited: time.Now(),
        	Score: 0.0,
    	}

		//cant find entry, create a new one
		if err = tag.GetTag(c, visitorId); err != nil {
			err = tag.StoreTag(c, visitorId)
		}
		if nil == err {
			offer = tag.CheckOffer()
		}	
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(offer))	
	}
}

/**
	Messenger handler
*/
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
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte(answer))
	} else {
		log.Infof(c,"Handling POST method")

		//This is a POST so expect FB Message
		body, err = ioutil.ReadAll(req.Body)
		if nil == err {
			log.Infof(c,"No error trying to parse body")
			bodyParts := new(Body)
			err = json.Unmarshal(body, bodyParts)
			if nil == err {
				log.Infof(c,"No error unmarshalling bodyparts")

				for _, messages := range bodyParts.Entry[0].Messaging {
					
					log.Infof(c,"Getting message")

					// Substitute words and get Intent with Entities
					message := new(nl.Message)
					message.Context = c
					_, messages.Message.Text = message.ReplaceWords(messages.Message.Text)
					err = message.GetIntent(messages.Message.Text)
					if err == nil && message.Entities.Intent != nil {
						log.Infof(c,"Getting answer for message " + message.Entities.Intent[0].Value)

						//Answer based on rules
						a := new(nl.Answer)
						a.Message = *message
						err, answer = a.GetAnswer()
						log.Infof(c,"Get ready to send answer " + answer)
						
						//send a resposne
						err = message.SendResponse(answer,messages.Sender.ID)
					}

				}

			}

		}

	}

	if nil != err {
			log.Infof(c,"Error:" + err.Error())
	}
	
	//Set for cross domain stuff
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Write([]byte(answer))
}

