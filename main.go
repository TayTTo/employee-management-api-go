package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)


var db *sql.DB

type employees struct {
	ID int64
	BirthDate string
	FirstName string
	LastName string
	Gender string
	HireDate string
}

func main() {
	cfg := mysql.NewConfig()
	cfg.User = os.Getenv("DBUSER")
	cfg.Passwd = os.Getenv("DBPASS")
	cfg.Net = "tcp"
	cfg.Addr = "127.0.0.1:3306"
	cfg.DBName = "employees"

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")
}

func getEmployeesByName(firstName string) ([]employees, error) {
	var employees []employees
	rows, err := db.Query("SELECT * FROM employees WHERE first_name = ?", firstName)
	if err != nil {
		return nil, fmt.Errorf("getEmployeesByName %q: %v", firstName, err);
	}
	
	for rows.Next() {
		var employ employees
		if err := rows.Scan(&employ.ID, &employ.BirthDate, &employ.FirstName)
	}
}


