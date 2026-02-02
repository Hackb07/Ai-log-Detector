package main

import (
	"ai-log-detector/internal/ingestion"
	"ai-log-detector/internal/mlclient"
	"ai-log-detector/internal/processor"
	"ai-log-detector/internal/storage"
	"log"
	"net/http"
)

func main() {
	// Configuration
	mlServiceURL := "http://localhost:8000"
	port := ":8080"

	// Ensure log directory exists
	// (Directories should have been created by setup, but good to ensure)

	// Initialize Dependencies
	store := storage.NewFileStore("logs/normal_logs.json", "logs/anomalous_logs.json")
	mlClient := mlclient.NewMLClient(mlServiceURL)

	// Initialize Processor
	proc := processor.NewLogProcessor(3, mlClient, store) // 3 workers
	proc.Start()
	defer proc.Stop()

	// Initialize Handler
	handler := ingestion.NewLogHandler(proc)

	// Routes
	http.Handle("/api/v1/logs", handler)

	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
