package main

import (
	"log"
	"math/rand"
	"time"

	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func jobDoStuff() {
	// Create a span for a web request at the /posts URL.
	span := tracer.StartSpan("job.doStuff", tracer.ResourceName("posts"))
	defer span.Finish()

	log.Println("Doing some work")
	// Do some work
	time.Sleep(200*time.Millisecond + time.Duration(rand.Intn(100))*time.Millisecond)
	log.Println("Done")

	// Set metadata
	span.SetTag("my_tag", "my_value")
}

func main() {
	// Start the tracer with zero or more options.
	tracer.Start(tracer.WithServiceName("my-app-worker"))
	defer tracer.Stop()

	for {
		jobDoStuff()
	}
}
