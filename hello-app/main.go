package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type response struct {
	Message string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := response{
		Message: "Hello world!",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode respoinse")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func byeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	res := response{
		Message: "Bye world!",
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Printf("Failed to encode respoinse")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func main() {

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/bye", byeHandler)

	fmt.Println("Server listend on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic("Failed to run server !")
	}

}
