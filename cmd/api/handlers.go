package main

import (
	"employee-management/internal/data"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func (app *application) getEmployeeByName(w http.ResponseWriter, r *http.Request) {
	var employees []data.Employees
	firstName := r.PathValue("name")

	if err := app.models.Movies.GET(firstName, &employees); err != nil {
		app.logger.Error(err.Error())
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

/* func (app *application) getEmployeeById(w http.ResponseWriter, r *http.Request) {
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
} */

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
