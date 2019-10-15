package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/tasi788/gocc"
)

// ConvertResponse make response to http request
type ConvertResponse struct {
	Status bool
	Text   string
}

func convert(args string, target string) (bool, string) {
	converter, err := gocc.New(target)
	if err != nil {
		log.Println(err)
		return false, "Reading Dictionary Failed!"
	}
	out, err := converter.Convert(args)
	if err != nil {
		log.Println(err)
		return false, "Conter Error!"
	}
	return true, out
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Method, r.URL)
	if r.Method == "get" {
		http.Redirect(w, r, "https://www.google.com", 301)
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}

	text := r.FormValue("text")
	target := r.FormValue("target")
	if len(text) == 0 || len(target) == 0 {
		loadResp := ConvertResponse{false, "Args Not Found "}
		resp, err := json.Marshal(loadResp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else {
			w.Write(resp)
		}
		return
	}
	status, result := convert(text, target)
	loadResp := ConvertResponse{status, result}
	resp, err := json.Marshal(loadResp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(resp)
	}

}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.HandleFunc("/", handler)
	log.Println(fmt.Sprintf("Server on http://localhost:%s", port))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
