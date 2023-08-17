package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	http.HandleFunc("/api/user", func(w http.ResponseWriter, request *http.Request) {
		b, _ := json.Marshal(User{
			Id:   "1",
			Name: "pomka",
		})
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, string(b))
	})
	http.ListenAndServe(":4000", nil)
}
