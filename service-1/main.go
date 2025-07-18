package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Message string `json:"message"`
}

func main() {
	log.Println("API1 starting on :8080")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("API1: Received request, calling API2")

		service2URL := os.Getenv("SERVICE_2_URL")
		if service2URL == "" {
			service2URL = "http://localhost:8081"
		}

		resp, err := http.Get(service2URL)
		if err != nil {
			log.Printf("API1: Error calling API2: %v", err)
			w.WriteHeader(500)
			json.NewEncoder(w).Encode(Response{Message: "Error calling API2"})
			return
		}
		defer resp.Body.Close()

		var api2Response Response
		json.NewDecoder(resp.Body).Decode(&api2Response)

		log.Printf("API1: Got response from API2: %s", api2Response.Message)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(api2Response)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
