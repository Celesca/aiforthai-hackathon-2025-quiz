package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type Request struct {
	Message string `json:"message"`
}

type Response struct {
	Message     string      `json:"message"`
	Status      string      `json:"status"`
	Timestamp   string      `json:"timestamp"`
	ProcessedBy string      `json:"processed_by"`
	FromAPI2    interface{} `json:"from_api2,omitempty"`
}

type ErrorResponse struct {
	Error     string `json:"error"`
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
}

func main() {
	// API1 - Golang Service
	fmt.Println("🚀 API1 (Golang) starting on port :8080")
	log.Println("API1: Ready to receive requests and forward to API2")

	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/api/hello", apiHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("API1: Received %s request to %s", r.Method, r.URL.Path)

	// Forward request to API2 (Python service)
	api2Response, err := forwardToAPI2("/")
	if err != nil {
		log.Printf("API1: Error communicating with API2: %v", err)
		sendErrorResponse(w, fmt.Sprintf("Failed to communicate with API2: %v", err), http.StatusInternalServerError)
		return
	}

	response := Response{
		Message:     "Hello from API1 (Golang)!",
		Status:      "success",
		Timestamp:   time.Now().Format(time.RFC3339),
		ProcessedBy: "API1-Golang",
		FromAPI2:    api2Response,
	}

	log.Printf("API1: Sending response back to user")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("API1: Received %s request to /api/hello", r.Method)

	var req Request
	if r.Method == "POST" && r.Body != nil {
		json.NewDecoder(r.Body).Decode(&req)
	}

	if req.Message == "" {
		req.Message = "Hello World from user!"
	}

	// Forward request to API2
	api2Response, err := forwardToAPI2("/api/hello")
	if err != nil {
		log.Printf("API1: Error communicating with API2: %v", err)
		sendErrorResponse(w, fmt.Sprintf("Failed to communicate with API2: %v", err), http.StatusInternalServerError)
		return
	}

	response := Response{
		Message:     fmt.Sprintf("API1 processed: %s", req.Message),
		Status:      "success",
		Timestamp:   time.Now().Format(time.RFC3339),
		ProcessedBy: "API1-Golang",
		FromAPI2:    api2Response,
	}

	log.Printf("API1: Successfully processed request and got response from API2")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func forwardToAPI2(endpoint string) (interface{}, error) {
	service2URL := os.Getenv("SERVICE_2_URL")
	if service2URL == "" {
		service2URL = "http://localhost:8081" // fallback for local development
	}

	url := service2URL + endpoint
	log.Printf("API1: Forwarding request to API2 at %s", url)

	// Create HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	// Send request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request to API2: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response from API2: %v", err)
	}

	// Parse JSON response
	var jsonResponse interface{}
	if err := json.Unmarshal(responseBody, &jsonResponse); err != nil {
		// If not JSON, return as string
		return string(responseBody), nil
	}

	log.Printf("API1: Successfully got response from API2")
	return jsonResponse, nil
}

func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	errorResp := ErrorResponse{
		Error:     message,
		Status:    "error",
		Timestamp: time.Now().Format(time.RFC3339),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(errorResp)
}
