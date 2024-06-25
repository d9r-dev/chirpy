package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleValidateChirpt(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		log.Printf("Error decoding parameters: %s", err)
	}

	type successBody struct {
		Valid bool `json:"valid"`
	}

	type errorBody struct {
		Error string `json:"error"`
	}

	w.Header().Set("Content-Type", "application/json")

	var dat []byte

	if len(params.Body) > 140 {

		respBody := errorBody{
			Error: "Chirp is too long",
		}

		dat, err = json.Marshal(respBody)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			w.WriteHeader(500)
			return
		}
		w.WriteHeader(400)
		w.Write(dat)
		return
	}

	respBody := successBody{
		Valid: true,
	}

	dat, err = json.Marshal(respBody)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
	}

	w.WriteHeader(200)
	w.Write(dat)
}
