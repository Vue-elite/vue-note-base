package server

import "net/http"

func InitializeMux() *http.ServeMux {
	MyMux := http.NewServeMux()
	setupRouter(MyMux)
	return MyMux
}
