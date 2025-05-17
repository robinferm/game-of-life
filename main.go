package main

import (
	"context"
	"net/http"
)

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}

func test(w http.ResponseWriter, req *http.Request) {
	component := grid()
	component.Render(context.Background(), w)
}
