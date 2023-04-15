package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const envWorkerUrl = "WORKER_URL"

var random = rand.New(rand.NewSource(time.Now().Unix()))

type WorkRequest struct {
	Id   int
	Name string
}

func main() {
	workerUrl := os.Getenv(envWorkerUrl)
	if workerUrl == "" {
		panic("worker url was missed")
	}
	mux := http.NewServeMux()

	client := &http.Client{Transport: &http.Transport{
		DisableKeepAlives: true,
	}}

	mux.Handle("/send-work", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		req := &WorkRequest{Id: random.Intn(1000), Name: fmt.Sprintf("work-%d", random.Intn(1000))}

		reqBytes, err := json.Marshal(req)
		if err != nil {
			log.Printf("can't marshall request: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		_, err = client.Post(workerUrl+"/work", "application/json", bytes.NewBuffer(reqBytes))

		if err != nil {
			log.Printf("error response from worker: %v\n", err)

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		log.Printf("sent work from: %s\n", r.Host)

		w.WriteHeader(http.StatusOK)
	}))

	s := &http.Server{Handler: mux, Addr: fmt.Sprintf(":%d", 8080)}

	log.Printf("starting server on %s", s.Addr)
	err := s.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("http server error: %v\n", err)
	}
}
