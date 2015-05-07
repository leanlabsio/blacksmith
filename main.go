package main

import (
	"gitlab.com/blacksmith/go-reactor/push"
	"net/http"
)

func main() {
	http.Handle("/push", &push.PushHandler{})
	http.ListenAndServe(":8080", nil)
}
