package main

import (
	"fmt"
	"log"
	"net/http"

	"azure-playground-generator/internal/config"
	"azure-playground-generator/pkg/api"
)

func main() {
	// Setup Config
	err := config.ParseEnvironment()
	if err != nil {
		log.Fatalf("Failed to parse env: %v", err)
	}

	// Setup Server
	mux := http.NewServeMux()
	mux.HandleFunc("/HttpTriggerWithOutputs", api.HTTPTriggerHandler)

	// Start Server
	fmt.Println("Go server Listening...on FUNCTIONS_CUSTOMHANDLER_PORT:", config.FunctionHTTPWorkerPort())
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.FunctionHTTPWorkerPort()), mux))
}
