package main

import (
	"github.com/cybersane/vue-note-base/handler/server"
	"net/http"
)

func main() {
	myMux := server.InitializeMux()
	http.ListenAndServe(":8080", myMux)
}