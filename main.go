package main

// # Copyright 2024 Google LLC

// # Licensed under the Apache License, Version 2.0 (the "License");
// # you may not use this file except in compliance with the License.
// # You may obtain a copy of the License at

// #     https://www.apache.org/licenses/LICENSE-2.0

// # Unless required by applicable law or agreed to in writing, software
// # distributed under the License is distributed on an "AS IS" BASIS,
// # WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// # See the License for the specific language governing permissions and
// # limitations under the License.

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/redis/go-redis/v9"
)

var redisPort string

func init() {
	redisPort = os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}
}

func main() {
	router := httprouter.New()
	router.GET("/health", healthCheck)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func healthCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	host := fmt.Sprintf("localhost:%s", redisPort)

	client := redis.NewClient(&redis.Options{
		Addr: host,
	})

	// Ping the local Redis instance.
	res, err := client.Ping(r.Context()).Result()
	if err != nil {
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
		log.Printf("Redis ping failed: %v", err)
		return
	}

	if res != "PONG" && res != "LOADING" && res != "MASTERDOWN" {
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
		log.Printf("Redis ping returned: %s", res)
	}

	// Potential to add additional custom checks here

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Redis health check successful: ", res)
}