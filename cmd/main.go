package main

import (
	"fmt"
	"net/http"
)

type User struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func main() {
	http.HandleFunc("/api/greeting", func(w http.ResponseWriter, request *http.Request) {
		b := fmt.Sprintf(`{"message": "Hello, %s from go"}`, request.URL.Query().Get("name"))
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, b)
	})
	http.ListenAndServe(":4000", nil)
}
