package crud

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID    int
	Name  string
	Email string
	Role  string
}

func DbConn(db_name string) (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "password"
	dbName := db_name
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

//Creating Table
func CreateTable(DBName string, tableName string) {
	db := DbConn(DBName)
	defer db.Close()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v(id int PRIMARY KEY AUTO_INCREMENT, Name varchar(30) NOT NULL, Email varchar(30), role varchar(30));", tableName)
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

//Create
func InsertData(db *sql.DB, name string, email string, role string) error {
	_, err := db.Exec("INSERT INTO employee (Name,Email,role) values(?,?,?)", name, email, role)
	if err != nil {
		return errors.New("not inserted")
	}
	return nil
}

// Read
func GetDetailsById(db *sql.DB, id int) (*Employee, error) {
	var empl Employee
	err := db.QueryRow("select * from employee where id=?", id).Scan(&empl.ID, &empl.Name, &empl.Email, &empl.Role)
	if err != nil {
		return nil, err
	}
	return &empl, nil
}

//Update
func UpdateById(db *sql.DB, id int, Name string, Email string, role string) error {

	res, err := db.Prepare("UPDATE employee SET name = ?, email = ?, role = ? WHERE id = ?")
	if err != nil {
		return errors.New("query doesn't prepare")
	}
	defer res.Close()
	_, err2 := res.Exec(Name, Email, role, id)

	if err2 != nil {
		return err2
	}
	return nil
}

//Delete
func DeleteById(db *sql.DB, id int) error {

	_, err := db.Exec("delete from employee where id=?", id)

	if err != nil {
		return err
	}

	return nil

}
