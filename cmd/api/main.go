package main 

import (
	"fmt"
	"log"
	"flag"
	"net/http"
	"time"
	"os"
	"database/sql"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env string
	dsn string //data name service or connection string to connect to database
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4001, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environment (dev, test, prod)")
    flag.StringVar(&cfg.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
	app := &application{
		config: cfg, 
		logger: logger,
	}
   // dsn := "user=postgres password=password dbname=readinglist sslmode=disable"
    db, err := sql.Open("postgres", cfg.dsn)
	//db, err := sql.Open("postgres", dsn)
	
	if err != nil {
		logger.Fatal(err)
	}

    defer db.Close()
    // test connection

    err = db.Ping()
    if err != nil {
		logger.Fatal(err)
	}
    
	logger.Printf("database connection pool established")

    
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

	errr := srv.ListenAndServe()
	logger.Fatal(errr)
}


// Middleware to log all incoming requests
func logRequest(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
		next(w, r)
	}
}
