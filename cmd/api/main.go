package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Application version
const version = "1.0.0"

// Configuration settings
type config struct {
	port int
	env  string
}

// Dependencies for HTTP handlers, helpers, and middleware.
type application struct {
	config config
	logger *log.Logger
}

func main() {
	// Instance of the config
	var cfg config

	// Read value from the command-line flags into config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize logger to write to the standard out stream with date and time
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Instance of the application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Declare HTTP server with some settings
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start the server
	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
