package data

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/go-sql-driver/mysql"
)

type Models struct {
	Movies MovieModel
}

func NewConnection() (*sql.DB, error) {
	var db *sql.DB
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "employees"
	
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, fmt.Errorf("Database connection error %q: %v", err, err)
	}
	return db, nil
}

func NewModel (db *sql.DB) Models {
	return Models {
		MovieModel {
			DB: db,
		},
	}
}

