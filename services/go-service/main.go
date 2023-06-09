package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello from Go! Hey, I'm live now!")
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running...")

	http.ListenAndServe(":80", nil)
}
