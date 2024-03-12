package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) { //to define http handler
	respondWithJSON(w, 200, struct{}{})
}
