package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RegisterUser struct {
	FirstName   string  `json:"firstName"`
	LastName string `json:"lastName"`
	EmailId string `json:"emailId"`
	Password string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}

func handleRegister(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var u RegisterUser
		if r.Body == nil {
			http.Error(w, "Please send a request body", 400)
			return
		}
		fmt.Println(r.Body)
		err := json.NewDecoder(r.Body).Decode(&u)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		} else {
			fmt.Println(u.FirstName)
			fmt.Println(u.LastName)
			fmt.Println(u.EmailId)
			fmt.Println(u.Password)
			fmt.Println(u.PasswordConfirmation)

			w.Write([]byte("<h1>Registeration page</h1>"))
		}
	}

}

func handleSignIn(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Sign In page</h1>"))
}

func getBoards(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Boards page</h1>"))
}
