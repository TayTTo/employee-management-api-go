package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
)


var db *sql.DB

type Employees struct {
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
	employee, err := getEmployeesByName("Sachin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Employee found: %v\n", employee);

	emp, err := getEmployeeById(490548)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Employee found: %v", emp)

	empID, err := addEmployee(Employees{
		ID: 999999,
		BirthDate: "2003-08-09",
		FirstName: "Anh",
		LastName: "Do",
		Gender: "M",
		HireDate: "2025-07-30",
	})
	
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added employee: %v\n", empID)
}

func getEmployeesByName(firstName string) ([]Employees, error) {
	var employees []Employees
	rows, err := db.Query("SELECT * FROM employees WHERE first_name = ?", firstName)
	if err != nil {
		return nil, fmt.Errorf("getEmployeesByName %q: %v", firstName, err);
	}
	
	defer rows.Close()
	for rows.Next() {
		var employee Employees
		if err := rows.Scan(&employee.ID, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate); err != nil {
			return nil, fmt.Errorf("getEmployeesByName %q: %v", firstName, err)
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", firstName, err)
	}
	return employees, nil
}

func getEmployeeById(employeeId int64) (Employees, error) {
	var employee Employees;
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
