package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/0xAckerMan/HNGx/0x01-people/cmd/data"
	_ "github.com/lib/pq"
)

var Version = "1.0.0"

type Config struct{
    Env string
    Port int
    dsn string
}

type Application struct{
    Config
    logger *log.Logger
    data.UserModel
}

func main() {
    var cfg Config

    connectionStr := "user=r00t password=password dbname=hngx2 port=5432 host=localhost sslmode=disable"

    flag.IntVar(&cfg.Port, "port", 3000, "The Webservice port")
    flag.StringVar(&cfg.Env, "env", "dev", "Environment(dev|staging|prod)")
    flag.StringVar(&cfg.dsn, "db-dsn", connectionStr, "Postgres DSN")
    flag.Parse()

    logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)
    app := Application{
        Config: cfg,
        logger: logger,

    }

    db, err := sql.Open("postgres", cfg.dsn)
    if err != nil{
        logger.Fatal(err)
    }

    defer db.Close()

    err = db.Ping()
    if err != nil{
        logger.Fatal(err)
    }
    
    logger.Println("Database connection pool established")

    addr := fmt.Sprintf(":%d", app.Port)
    logger.Printf("Starting %s on port %s", app.Env, addr)
    srv := http.Server{
        Addr: addr,
        Handler: app.routes(),
        IdleTimeout: time.Minute,
        ReadTimeout: 20 * time.Second,
        WriteTimeout: 20 * time.Second,
    }
    err =srv.ListenAndServe()
    if err != nil {
        fmt.Println(err)
    }
}

