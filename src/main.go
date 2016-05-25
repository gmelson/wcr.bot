// wcr.bot project main.go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func ProcessRequest(w http.ResponseWriter, req *http.Request) {
	//message := req.FormValue("message")

	//Set for cross domain stuff
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	http.HandleFunc("/webhook", ProcessRequest)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error exiting")
	}

	fmt.Println("Listening on port 8080")

}
