package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	pb "./irisc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	// Declare command-line flags
	debug := flag.Int("debug", -1, "Run client in debug mode. 0 for summary, 1 for request, 2 for response.")
	continueOnError := flag.Bool("continueOnError", false, "Print warning instead of crashing")
	address := flag.String("address", "", "Adress:Port to be used for client")
	requests := flag.Int("requests", 10, "Number of total requests to be done.")
	concurrency := flag.Int("concurrency", 1, "Number of workers in parallel")
	flag.Parse()

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(*address, opts...)
	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewPredictorClient(conn)

	sem := make(chan int, *concurrency)
	var wg sync.WaitGroup
	start := time.Now()
	for i := 0; i < *requests; i++ {
		// will block if there is concurrency ints in sem
		sem <- 1
		// Increment the WaitGroup counte
		wg.Add(1)
		go func(i int) {
			// removes an int from sem, allowing another to proceed
			defer func() {
				<-sem
				wg.Done()
			}()

			request := &pb.IrisRequest{
				PetalLength: rand.Float64(),
				PetalWidth:  rand.Float64(),
				SepalLength: rand.Float64(),
				SepalWidth:  rand.Float64(),
			}

			if *debug == 0 {
				fmt.Println("Sending request for", i)
			}
			if *debug == 1 {
				fmt.Printf("Request : %v\n", request)
			}
			response, err := client.Predict(context.Background(), request)
			if err != nil {
				if *continueOnError {
					grpclog.Errorf("Request failed")
				} else {
					grpclog.Fatalf("Request failed")
				}
			}

			if *debug == 2 {
				fmt.Printf("Response : %v\n", response.Species)
			}

			if *debug == 0 {
				fmt.Println("Received response for", i)
			}
		}(i)
	}
	// ensure everything is computed
	wg.Wait()
	close(sem)
	elapsed := time.Since(start)
	log.Printf("Took %s to compute %d loops", elapsed, *requests)
}
