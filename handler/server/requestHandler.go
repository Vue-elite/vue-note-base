package server

import "net/http"

func handleRegister(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Registeration page</h1>"))
}

func handleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Sign In page</h1>"))
}

func getBoards(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Boards page</h1>"))
}
