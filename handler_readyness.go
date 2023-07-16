package main

import "net/http"

func handlerReadyness(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, struct{}{})
}
