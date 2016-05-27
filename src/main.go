// wcr.bot project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"nl"
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

func ProcessRequest(w http.ResponseWriter, req *http.Request) {
	q := req.FormValue("message")
	var answer string

	//Handle messenger challenge if GET otherwise handle message
	if req.Method == "GET" {
		if req.URL.Query().Get("hub.verify_token") == "<validation_token>" {
			answer = req.URL.Query().Get("hub.challenge")
		} else {
			answer = "Error, wrong validation token"
		}
	} else {
		//This is a POST so expect FB Message
		req.Body.Read()
		// Substitute words and get Intent with Entities
		message := new(nl.Message)
		_, q = message.ReplaceWords(q)
		err := message.GetIntent(q)
		if err == nil {
			//Answer based on rules
			a := new(nl.Answer)
			a.Message = *message
			err, answer = a.GetAnswer()
		}

	}

	//Set for cross domain stuff
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Write([]byte(answer))
}

func main() {
	http.HandleFunc("/webhook", ProcessRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error exiting")
	}

	fmt.Println("Listening on port 8080")

}
