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

	for {
		req := &WorkRequest{Id: random.Intn(1000), Name: fmt.Sprintf("work-%d", random.Intn(1000))}

		reqBytes, err := json.Marshal(req)
		if err != nil {
			log.Printf("can't marshall request: %v", err)
			continue
		}

		_, err = http.Post(workerUrl+"/work", "application/json", bytes.NewBuffer(reqBytes))

		if err != nil {
			log.Printf("error response from worker: %v\n", err)
			continue
		}

		log.Printf("sent work: %+v\n", req)
	}
}
