package main

import (
	"log"
	"net/http"
)

type apiConfig struct {
	fileserverHits int
}

func main() {
	const port = "8080"
	apiCfg := apiConfig{
		fileserverHits: 0,
	}
	mux := http.NewServeMux()
	mux.Handle("/app/*", apiCfg.middlewarMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(".")))))
	mux.HandleFunc("GET /api/healthz", handlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("GET /api/reset", apiCfg.handlerReset)
	server := &http.Server{Addr: ":" + port, Handler: mux,}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
