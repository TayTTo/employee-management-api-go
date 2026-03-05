package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

const version = "1.0.0"

var db *sql.DB

type config struct {
	port int
	env  string
}

type application struct {
	config config
	logger *slog.Logger
}

type Employees struct {
	ID        int64
	BirthDate string
	FirstName string
	LastName  string
	Gender    string
	HireDate  string
}

func main() {
	var cfg config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()
	dbcfg := mysql.NewConfig()
	dbcfg.User = os.Getenv("DBUSER")
	dbcfg.Passwd = os.Getenv("DBPASS")
	dbcfg.Net = "tcp"
	dbcfg.Addr = "127.0.0.1:3306"
	dbcfg.DBName = "employees"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		config: cfg,
		logger: logger,
	}

	var err error
	db, err = sql.Open("mysql", dbcfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	/* pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!") */
	svr := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog: slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}
	logger.Info("starting server", "addr", svr.Addr, "env", cfg.env)
	err = svr.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

