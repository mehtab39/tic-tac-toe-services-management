package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
)

const scriptDir = "scripts/"

func main() {
	// Define a file server to serve static files
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Handle requests to start services
	http.HandleFunc("/start/", handleStartService)

	// Handle requests to stop services
	http.HandleFunc("/stop/", handleStopService)

	// Handle requests to check health
	http.HandleFunc("/health/", handleCheckHealth)

	// Start the HTTP server and listen for requests on port 1234
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		panic(err)
	}
}

func extractIDFromURL(path string) string {
	// Extract the ID from the URL path
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func handleStartService(w http.ResponseWriter, r *http.Request) {
	serviceID := extractIDFromURL(r.URL.Path)
	fmt.Printf("Starting service with ID: %s\n", serviceID)
	cmd, _ := executeCmd(serviceID, "start")
	standardOutput(w, cmd)
}

func handleStopService(w http.ResponseWriter, r *http.Request) {
	serviceID := extractIDFromURL(r.URL.Path)
	fmt.Printf("Stopping service with ID: %s\n", serviceID)
	cmd, _ := executeCmd(serviceID, "stop")
	standardOutput(w, cmd)
}

func handleCheckHealth(w http.ResponseWriter, r *http.Request) {
	serviceID := extractIDFromURL(r.URL.Path)
	fmt.Printf("Checking health of service with ID: %s\n", serviceID)
	cmd, _ := executeCmd(serviceID, "health")
	standardOutput(w, cmd)
}

func executeCmd(serviceID string, taskId string) (*exec.Cmd, error) {
	scriptPath := scriptDir + serviceID + "/" + taskId + ".sh"

	fmt.Printf("executing command %s\n", scriptPath)
	cmd := exec.Command("/bin/bash", scriptPath)
	return cmd, nil
}

func standardOutput(w http.ResponseWriter, cmd *exec.Cmd) {
	// Capture standard output
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error executing shell script: %s\n", err)
		http.Error(w, "Error executing shell script", http.StatusInternalServerError)
		return
	}

	// Write output to the HTTP response
	w.Header().Set("Content-Type", "text/plain")
	w.Write(output)
}
