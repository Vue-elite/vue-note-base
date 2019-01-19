package server

import "net/http"

func setupRouter(myMux *http.ServeMux) {
	myMux.HandleFunc("/", handleRegister)
	myMux.HandleFunc("/register", handleRegister)
	myMux.HandleFunc("/sign-in", handleSignIn)
	myMux.HandleFunc("/boards", getBoards)
}