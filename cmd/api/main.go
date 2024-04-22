package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config // instantiate config
	// define flags using stdlib's flag package
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.IntVar(&cfg.port, "p", 4000, "API server port (shorthand)")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev|stage|prod)")
	flag.StringVar(&cfg.env, "e", "dev", "Environment (dev|stage|prod) (shorthand)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime) // instantiate logger

	app := &application{ // instantiate app with our logger and config
		config: cfg,
		logger: logger,
	}

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{ // instantiate http server from stdlib's net/http packagef
		Addr:         addr, // configure server stuff
		Handler:      app.route(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("Starting %s server on %s", cfg.env, addr)
	err := srv.ListenAndServe() // start the server (err captures any error returned)
	logger.Fatal(err)           // any err thrown will be logged here
}
