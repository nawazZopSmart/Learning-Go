package main

import (
	"fmt"
	"log"

	s "github.com/SN786/sqlpractise/crud"
)

func main() {
	DBName := "emp"
	tableName := "employee"
	s.CreateTable(DBName, tableName)
	DbCon := s.DbConn("emp")

	err := s.InsertData(DbCon, "Ram Vir", "rv@r.com", "Intern")
	erri := s.InsertData(DbCon, "Richesh", "rich@r.com", "Intern")

	if err != nil {
		log.Fatalf("Inser Err: %v", err)
	}
	if erri != nil {
		log.Fatalf("Inser Err: %v", err)
	}
	// var empDet Employee;
	empDet, err := s.GetDetailsById(DbCon, 1)

	if err != nil {
		log.Fatalf("GetDetailsById Err: %v", err)
	}

	fmt.Println(empDet.Name, empDet.Email, empDet.Role)

	err2 := s.UpdateById(DbCon, 1, "Sharif", "k@k.com", "SDE-II")
	if err2 != nil {
		log.Fatalf("GetUpdateById Err: %v", err)
	}
	err3 := s.DeleteById(DbCon, 1)
	if err3 != nil {
		log.Fatalf("DeleteById Err: %v", err)
	}

}
