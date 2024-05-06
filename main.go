package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type Response struct {
	// your xml response
}

func xmlHandler(w http.ResponseWriter, r *http.Request) {

	// check the http is POST
	if r.Method != http.MethodPost {
		http.Error(w, "http needed POST method", http.StatusMethodNotAllowed)
		return
	}

	// streaming r.Body content
	StreamReader(r.Body)
	defer r.Body.Close()

	// build xml response
	res := Response{
		// your xml response
	}

	// marshal to xml format
	response, err := xml.MarshalIndent(res, "", "  ")
	if err != nil {
		http.Error(w, "Error marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")

	w.Write(response)
}

func main() {
	http.HandleFunc("/parser", xmlHandler)
	fmt.Println("server running")
	http.ListenAndServe(":8080", nil)
}
