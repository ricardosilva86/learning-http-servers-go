package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	apiCfg := apiConfig{}
	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	mux.Handle("/logo.png", http.FileServer(http.Dir("./assets")))
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /admin/metrics", apiCfg.middlewareGetMetrics)
	mux.HandleFunc("POST /admin/reset", apiCfg.middlewareReset)
	mux.HandleFunc("GET /api/healthz", handleHealthz)
	mux.HandleFunc("/api/validate_chirp", handleValidateChirp)

	log.Printf("starting server at %s\n", server.Addr)
	log.Fatal(server.ListenAndServe())
}

func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, err := io.WriteString(w, "OK")
	if err != nil {
		fmt.Printf("error writing to response writer: %w\n", err)
		return
	}
}
