package data

import (
	"database/sql"
	"fmt"
)


type Employees struct {
	ID        int64
	BirthDate string
	FirstName string
	LastName  string
	Gender    string
	HireDate  string
}

type MovieModel struct {
	DB *sql.DB
}

func (m MovieModel) GET (firstName string, empList *[]Employees) (error) {
	rows, err := m.DB.Query("SELECT * FROM employees WHERE first_name = ?", firstName)
	if err != nil {
		return fmt.Errorf("Insert query error %q: %v", err, err)
	}

	defer rows.Close()
	for rows.Next() {
		var employee Employees
		if err := rows.Scan(&employee.ID, &employee.BirthDate, &employee.FirstName, &employee.LastName, &employee.Gender, &employee.HireDate); err != nil {
			return fmt.Errorf("Insert query error %q: %v", err, err)
		}
		*empList = append(*empList, employee)
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("Error in fetching rows %q: %v", err, err)
	}
	return nil
}
