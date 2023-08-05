package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError( w http.ResponseWriter, code int, msg string){
	if code > 499{
		log.Println("reponding with 5XX error: %s", msg)
		
	}

	type errResponse struct {
		Error string `json:"error"`
	}
	respondWithJson(w,code, errResponse{
		Error: msg,
	})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to Marshall to JSON response: %v ", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}