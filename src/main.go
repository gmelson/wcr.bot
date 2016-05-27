// wcr.bot project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"nl"
)

func ProcessRequest(w http.ResponseWriter, req *http.Request) {
	q := req.FormValue("message")
	var answer string

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
