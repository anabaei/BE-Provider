package main 

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"time"
	"os"
	"database/sql"
	"github.com/lib/pq"
)

type config struct {
	port int
	env string
}


type application struct {
	config config
	logger *log.Logger
}


func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4001, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev, test, prod)")
    flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
	app := &application{
		config: cfg, 
		logger: logger,
	}

	// mux := http.NewServeMux()

	// Add router debug logging
	
    fmt.Println("Server starting on port 4001...")

	addr := fmt.Sprintf(":%d", cfg.port)

	srv := &http.Server{
		Addr: addr,
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	err := srv.ListenAndServe()
	logger.Fatal(err)
}


// Middleware to log all incoming requests
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
		next(w, r)
	}
}
