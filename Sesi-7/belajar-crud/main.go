package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "db_go_sql"
)

var (
	db  *sql.DB
	err error
)

type Employee struct {
	ID       int
	FullName string
	Email    string
	Age      int
	Division string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")

	// CreateEmployee()
	GetEmployees()
	GetEmployee(3)
	// UpdateEmployee()
	// DeleteEmployee()
}

func CreateEmployee() {
	var employee = Employee{}

	sqlStatement := `
		INSERT INTO employees (
			full_name, 
			email, age, 
			division
		) VALUES ($1, $2, $3, $4)
		Returning *
	`

	err = db.QueryRow(sqlStatement, "Khairul", "diah@gmail.com", 26, "Backend Developer").
		Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Email,
			&employee.Age,
			&employee.Division,
		)

	if err != nil {
		panic(err)
	}

	fmt.Printf("New Employee Data : %+v\n", employee)
}

func GetEmployees() {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees ORDER BY id ASC`
	rows, err := db.Query(sqlStatement)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Email,
			&employee.Age,
			&employee.Division,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee datas:", results)
}

func GetEmployee(id uint) {
	var results = []Employee{}

	sqlStatement := `SELECT * FROM employees WHERE id = $1 ORDER BY id ASC`
	rows, err := db.Query(sqlStatement, id)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	for rows.Next() {
		var employee = Employee{}

		err = rows.Scan(
			&employee.ID,
			&employee.FullName,
			&employee.Email,
			&employee.Age,
			&employee.Division,
		)

		if err != nil {
			panic(err)
		}

		results = append(results, employee)
	}

	fmt.Println("Employee data:", results)
}

func UpdateEmployee() {
	sqlStatement := `
		UPDATE 
			employees
		SET 
			full_name= $2, 
			email= $3, 
			division= $4, 
			age= $5
		WHERE 
			id = $1;`

	res, err := db.Exec(sqlStatement, 1, "Khairul", "kabdi3844@gmail.com", "Backend Dev", 26)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Updated data amount:", count)
}

func DeleteEmployee() {
	sqlStatement := `
	DELETE FROM 
		employees 
	WHERE 
		id = $1;
	`

	res, err := db.Exec(sqlStatement, 1)
	if err != nil {
		panic(err)
	}

	count, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println("Deleted data amount:", count)
}
