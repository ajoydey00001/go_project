package main

import (
	
	"net/http"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	// log.Println("Readiness probe")
	respondWithJSON(w,200,struct{}{})
}