package main

import (
	s "./crud/crud.go"
)

func main() {
	// db_name := "Employee_Db"
	table_name := "Employee_Details"
	// s.CreateTable(db_name, table_name)

	db_conn := s.DbConn("Employee_Db")
	// name := []string{"Richesh", "Sharif", "Sukant", "Ishan"}
	// email := []string{"r@r.com", "s@s.com", "ss@ss.com", "i@i.com"}
	// role := []string{"SDE-I", "SDE-I", "SDE-I", "SDE-I"}
	// err := s.InsertData(table_name, db_conn, name, email, role)

	// if err != nil {
	// 	log.Fatal("Insert: Error")
	// }

	// s.GetDetailsById(db_conn, 1, table_name)
	// s.GetDetailsById(db_conn, 2, table_name)
	// s.GetDetailsById(db_conn, 3, table_name)
	// s.UpdateById(db_conn, 4, "KAKA", "k@k.com", "SDE-II", table_name)
	// s.DeleteById(db_conn, 4)
	s.GetAll(db_conn, table_name)
}
