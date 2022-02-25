package crud

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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

func RecordExists(db *sql.DB, id int, table_name string) bool {
	res, err := db.Prepare("select id from Employee_Details where id = ?")
	if err != nil {
		log.Fatal(err)
	}
	_, err2 := res.Query(id)

	if err2 != nil {
		if err2 == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err2)
		}
	}
	return true
}

// func CreateTable()

func CreateTable(db_name string, table_name string) {
	db := DbConn("Employee_Db")
	defer db.Close()
	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %v(id int PRIMARY KEY AUTO_INCREMENT, Name varchar(30) NOT NULL, Email varchar(30), role varchar(30));", table_name)
	fmt.Println(query)
	res, err := db.Exec(query)
	if err != nil {
		fmt.Println("here")
		log.Fatal(err)
	}
	fmt.Println(res.RowsAffected())
}

// Create
func InsertData(table_name string, db *sql.DB, name []string, email []string, role []string) error {
	query := "INSERT INTO Employee_Details (Name,Email,role) values (?,?,?);"

	res, err := db.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(name); i++ {
		_, err2 := res.Exec(name[i], email[i], role[i])

		if err2 != nil {
			fmt.Println("here")
			log.Fatal(err2.Error())
			return err2
		}
		fmt.Println("Successfully Inserted!")
	}
	return nil
}

// Read
func GetDetailsById(db *sql.DB, id int, table_name string) {

	res, err := db.Prepare("select * from Employee_Details where id=?")

	if err != nil && len(table_name) != 0 {
		log.Fatal(err)
	}

	defer res.Close()

	var (
		id_val int
		Name   string
		Email  string
		role   string
	)

	err2 := res.QueryRow(id).Scan(&id_val, &Name, &Email, &role)
	if err2 != nil {
		if err2 == sql.ErrNoRows {
			fmt.Println("No Rows")
		} else {
			log.Fatal(err2)
		}
	}

	if id_val != 0 {
		fmt.Println("Name : ", Name, " Email:", Email, " Role : ", role)
	} else {
		fmt.Println("Read: Invalid Id")
	}

}

// Update
func UpdateById(db *sql.DB, id int, Name string, Email string, role string, table_name string) {

	res, err := db.Prepare("update Employee_Details set Name=?, Email=?, role=? where id=?")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Close()
	r2, err2 := res.Exec(Name, Email, role, id)

	rownum, _ := r2.RowsAffected()

	if rownum == 0 {
		log.Fatal("Update: Invalid Id, No record with this id.")
	}
	if err2 != nil {

		if err2 == sql.ErrNoRows {
			fmt.Println("Invalid Id")
		} else {
			log.Fatal(err2)
		}
		fmt.Print("Error : ", err2.Error())
	}
	fmt.Println("New Details are - ")
	GetDetailsById(db, id, table_name)
}

// Delete
func DeleteById(db *sql.DB, id int) {

	res, err := db.Prepare("delete from Employee_Details where id=?")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Close()
	res2, _ := res.Exec(id)

	if num, _ := res2.RowsAffected(); num == 0 {
		log.Fatal("Delete : Invalid Id")
	}
	fmt.Println("Successfully Deleted ", id)
}

// func GetAll(db *sql.DB,table_name string){

// 	query:="Select * from ?"
// 	res,err:=db.Prepare(query)
// 	if err!=nil{
// 		log.Fatal(err)
// 	}
// 	resp:=res.QueryRow(table_name)
// 	for
// }
