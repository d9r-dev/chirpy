package main

import (
	"log"
	"net/http"

	handlers "github.com/d9r-dev/chirpy/handlers"
	database "github.com/d9r-dev/chirpy/internals"
)

func main() {
	const filepathRoot = "."
	const port = "8080"

	db, err := database.NewDB("databse.json")
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := handlers.ApiConfig{
		FileserverHits: 0,
		DB:             db,
	}
	mux := http.NewServeMux()
	mux.Handle("/app/*", apiCfg.MiddlewarMetricsInc(http.StripPrefix("/app/", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("GET /api/healthz", handlers.HandlerReadiness)
	mux.HandleFunc("GET /admin/metrics", apiCfg.HandlerMetrics)
	mux.HandleFunc("GET /api/reset", apiCfg.HandlerReset)
	mux.HandleFunc("POST /api/chirp", apiCfg.HandleChirpsCreate)
	mux.HandleFunc("GET /api/chirps", apiCfg.HandlerChirpsRetrieve)
	server := &http.Server{Addr: ":" + port, Handler: mux}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}
