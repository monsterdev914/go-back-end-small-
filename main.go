package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	hello()
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Test endpoint")
}
func main() {
	InitDB()
	defer CloseDB()

	http.HandleFunc("/", greet)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":3001", nil)
}
