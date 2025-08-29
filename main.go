package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Payload struct {
	Name  string `json: "name`
	Email string `json: email`
}

type ResponseData struct {
	Message string `json:"message`
	Status  int    `json: status`
}

func greet(w http.ResponseWriter, r *http.Request) {
	hello()
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Test endpoint")
}

func postMethodTest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method is not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	var p Payload
	err = json.Unmarshal(body, &p)

	if err != nil {
		http.Error(w, "Failed to parse JSON!", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()
	fmt.Printf("Parsed Payload: %+v\n", p)
	fmt.Println("Recived Post data", string(body))

	data := ResponseData{
		Message: "Hello",
		Status:  200,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	// fmt.Fprintf(w, "Received: %+v", p)

}
func main() {
	InitDB()
	defer CloseDB()

	http.HandleFunc("/", greet)
	http.HandleFunc("/postTest", postMethodTest)
	http.HandleFunc("/test", test)
	http.ListenAndServe(":3001", nil)
}
