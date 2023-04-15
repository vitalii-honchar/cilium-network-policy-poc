package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type WorkRequest struct {
	Id   int
	Name string
}

func main() {

	mux := http.NewServeMux()
	mux.Handle("/work", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		var req WorkRequest

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			log.Printf("can't decode request: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("start work: %+v\n", req)
		w.WriteHeader(http.StatusOK)
	}))

	s := &http.Server{Handler: mux, Addr: fmt.Sprintf(":%d", 8081)}

	log.Printf("starting worker on %s", s.Addr)
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("http server error: %v\n", err)
	}
}
