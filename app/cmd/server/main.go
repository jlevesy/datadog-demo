package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func handleHi(w http.ResponseWriter, r *http.Request) {
	// Do some work
	time.Sleep(200*time.Millisecond + time.Duration(rand.Intn(100))*time.Millisecond)

	// Fail sometimes.
	switch v := rand.Intn(100); {
	case v > 95:
		w.WriteHeader(500)
		return
	case v > 85:
		w.WriteHeader(400)
		return
	}

	w.Write([]byte("hi\n"))
}

func main() {
	tracer.Start()
	defer tracer.Stop()
	mux := muxtrace.NewRouter(muxtrace.WithServiceName("test-app"))
	mux.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
