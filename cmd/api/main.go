package main

import (
	"employee-management/internal/data"
	"flag"
	"fmt"
	"log/slog"
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
	logger *slog.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	db, err := data.NewConnection()

	if err != nil {
		logger.Error(err.Error())
	}

	defer db.Close()

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModel(db),
	}

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
