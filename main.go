package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Employees struct {
	ID        int64
	BirthDate string
	FirstName string
	LastName  string
	Gender    string
	HireDate  string
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
	mux := http.NewServeMux()
	mux.HandleFunc("GET /getByName/{name}", getEmployeeByName)
	svr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	svr.ListenAndServe()

	// emp, err := getEmployeeById(490548)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Employee found: %v", emp)

	// empID, err := addEmployee(Employees{
	// 	ID: 999999,
	// 	BirthDate: "2003-08-09",
	// 	FirstName: "Anh",
	// 	LastName: "Do",
	// 	Gender: "M",
	// 	HireDate: "2025-07-30",
	// })

}

func getEmployeeByName(w http.ResponseWriter, r *http.Request) {
	var employees []Employees
	firstName := r.PathValue("name")
	rows, err := db.Query("SELECT * FROM employees WHERE first_name = ?", firstName)
	if err != nil {
		fmt.Errorf("getEmployeesByName %q: %v", firstName, err)
	}

	defer rows.Close()
	for rows.Next() {
		var employee Employees
		if err := rows.Scan(&employee.ID, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate); err != nil {
			fmt.Errorf("getEmployeesByName %q: %v", firstName, err)
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		fmt.Errorf("albumsByArtist %q: %v", firstName, err)
	}
	for _, emp := range employees {
		fmt.Println(emp)
		employees_json, err := json.MarshalIndent(employees, "", "	")
		if err != nil {
			log.Fatal(err)
		}
		first_name := []byte(employees_json)
		w.Write(first_name)
	}
}

func getEmployeeById(employeeId int64) (Employees, error) {
	var employee Employees
	row := db.QueryRow("SELECT * FROM employees WHERE emp_no = ?", employeeId)
	if err := row.Scan(&employee.ID, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate); err != nil {
		if err == sql.ErrNoRows {
			return employee, fmt.Errorf("getEmployeeById %d: no such album", employeeId)
		}
		return employee, fmt.Errorf("getEmployeeById %q: %v", employeeId, err)
	}
	return employee, nil
}

func addEmployee(employee Employees) (int64, error) {
	result, err := db.Exec("INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date) VALUES (?, ?, ?, ?, ?, ?)", employee.ID, employee.BirthDate, employee.FirstName, employee.LastName, employee.Gender, employee.HireDate)
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	return id, nil
}
