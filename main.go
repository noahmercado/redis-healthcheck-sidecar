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
var redisPwd string
var probePort string
var client *redis.Client

func init() {
	probePort = os.Getenv("PROBE_PORT")
	if probePort == "" {
		probePort = "8080"
	}

	redisPort = os.Getenv("REDIS_PORT")
	if redisPort == "" {
		redisPort = "6379"
	}

	redisPwd = os.Getenv("REDIS_PWD")
	if redisPwd == "" {
		fmt.Println("No password set for Redis")
	}

	host := fmt.Sprintf("localhost:%s", redisPort)

	client = redis.NewClient(&redis.Options{
		Addr: host,
		Password: redisPwd,
	})
}

func main() {
	router := httprouter.New()
	router.GET("/liveness", livenessCheck)
	router.GET("/readiness", readinessCheck)
	router.GET("/startup", startupCheck)

	port := fmt.Sprintf(":%s", probePort)
	log.Fatal(http.ListenAndServe(port, router))
}

func livenessCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("livenessCheck...")
	// Ping the local Redis instance.
	res, err := client.Ping(r.Context()).Result()
	if err != nil {
		log.Printf("Redis ping failed: %v", err)
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
		return
	}

	if res != "PONG" && res != "LOADING" && res != "MASTERDOWN" {
		log.Printf("Redis ping returned: %s", res)
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
	}

	// Potential to add additional custom checks here

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Redis health check successful: ", res)
}

func readinessCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Println("readinessCheck...")
	// Ping the local Redis instance.
	res, err := client.Ping(r.Context()).Result()
	if err != nil {
		log.Printf("Redis ping failed: %v", err)
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
		return
	}

	if res != "PONG" {
		log.Printf("Redis ping returned: %s", res)
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
	}

	// Potential to add additional custom checks here

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Redis health check successful: ", res)
}

func startupCheck(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	fmt.Println("startUpCheck...")
	// Ping the local Redis instance.
	res, err := client.Ping(r.Context()).Result()
	if err != nil {
		log.Printf("Redis ping failed: %v", err)
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
		return
	}

	if res != "PONG" && res != "LOADING" && res != "MASTERDOWN" {
		log.Printf("Redis ping returned: %s", res)
		http.Error(w, "Redis health check failed", http.StatusInternalServerError)
	}

	// Potential to add additional custom checks here

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Redis health check successful: ", res)
}