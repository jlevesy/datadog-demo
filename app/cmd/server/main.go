package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	httptrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/net/http"
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
	mux := httptrace.NewServeMux(httptrace.WithServiceName("test-app"))
	mux.HandleFunc("/hi", handleHi)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
