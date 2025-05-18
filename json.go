package main

import (
	"log"
	"net/http"

	"encoding/json"
)

func respondWithJSON(w http.ResponseWriter,code int,payload interface{}){
	dat , err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON response : %v\n", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

	
}

func respondWithError(w http.ResponseWriter,code int,message string){
	if(code > 499){
		log.Printf("Error %d: %s\n", code, message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	
	respondWithJSON(w, code, errorResponse{
		Error: message,
		
	})
}	
