package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func (app *application) getEmployeeByName(w http.ResponseWriter, r *http.Request) {
	var employees []Employees
	firstName := r.PathValue("name")
	rows, err := db.Query("SELECT * FROM employees WHERE first_name = ?", firstName)
	if err != nil {
		app.logger.Error(err.Error())
		os.Exit(1)
	}

	defer rows.Close()
	for rows.Next() {
		var employee Employees
		if err := rows.Scan(&employee.ID, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate); err != nil {
			log.Fatal(err)
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
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

func (app *application) getEmployeeById(w http.ResponseWriter, r *http.Request) {
	var employee Employees
	employeeId := r.PathValue("id")
	row := db.QueryRow("SELECT * FROM employees WHERE emp_no = ?", employeeId)
	if err := row.Scan(&employee.ID, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate); err != nil {
		if err == sql.ErrNoRows {
			app.logger.Error(err.Error())
		}
		app.logger.Error(err.Error())
	}
	employees_json, err := json.MarshalIndent(employee, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	employee_data := []byte(employees_json)
	w.Write(employee_data)
}

/* func (app *application) addEmployee(w http.ResponseWriter, r *http.Request) {
	result, err := db.Exec("INSERT INTO employees (emp_no, birth_date, first_name, last_name, gender, hire_date) VALUES (?, ?, ?, ?, ?, ?)", employee.ID, employee.BirthDate, employee.FirstName, employee.LastName, employee.Gender, employee.HireDate)
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addEmployee: %v", err)
	}
	return id, nil
} */
