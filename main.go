package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/yudafatah/product-api-learngolang/handlers"
)

func main() {

	port := os.Getenv("PORT")

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	s := &http.Server{
		Handler:           sm, // set the default handler
		ReadTimeout:       1 * time.Second, // max time to read request from the client
		ReadHeaderTimeout: 0,
		WriteTimeout:      1 * time.Second, // max time to write response to the client
		IdleTimeout:       120 * time.Second, // max time for connections using TCP Keep-Alive
		MaxHeaderBytes:    0,
		ErrorLog: l, // set the logger for the server
	}

	// start the server which already created
	go func() {
		l.Printf("Starting server on port %s\n", port)

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap signal term or interupt and gracefully shutdown the server
	SigChan := make(chan os.Signal)
	signal.Notify(SigChan, os.Interrupt)
	signal.Notify(SigChan, os.Kill)

	// Block until a signal is received.
	sig := <- SigChan
	l.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(),30*time.Second)
	s.Shutdown(ctx)

}