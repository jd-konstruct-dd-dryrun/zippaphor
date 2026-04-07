package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Status  string `json:"status"`
	Version string `json:"version,omitempty"`
	Name    string `json:"name,omitempty"`
}

func main() {
	mux := http.NewServeMux()

	version := os.Getenv("CHART_VERSION")
	if version == "" {
		version = "unknown"
	}

	image := os.Getenv("IMAGE")
	log.Printf("zippaphor starting version=%s image=%s", version, image)

	// Return JSON response with status and version
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Status:  "success",
			Version: version,
		})
	})

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		name := r.URL.Query().Get("name")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Response{
			Status: "success",
			Name:   name,
		})
	})

	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Fatal(srv.ListenAndServe())
}
