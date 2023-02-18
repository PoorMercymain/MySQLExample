package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type row struct {
	id    int64
	name  string
	value int
}

func addRow(db *sql.DB, name string, value int) (int64, error) {
	if len(name) > 20 {
		log.Fatal("Name cannot be longer than 20 characters")
	}

	result, err := db.Exec("INSERT INTO item (name, value) VALUES (?, ?);", name, value)
	if err != nil {
		log.Fatal("Error inserting row:", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Fatal("Error getting id of inserted row:", err)
	}

	return id, nil
}

func selectAllRows(db *sql.DB) []row {
	result, err := db.Query("SELECT * FROM item;")
	if err != nil {
		log.Fatal("Error getting rows:", err)
	}

	rows := make([]row, 0)
	var currentRow row

	for result.Next() {
		err = result.Scan(&currentRow.id, &currentRow.name, &currentRow.value)
		if err != nil {
			log.Fatal("Error scanning result rows:", err)
		}

		rows = append(rows, currentRow)
	}

	return rows
}

func main() {
	cfg := mysql.Config{
		User:   os.Getenv("MYSQL_USER"),
		Passwd: os.Getenv("MYSQL_PASSWORD"),
		DBName: os.Getenv("MYSQL_DATABASE"),
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal("Error while trying to connect to database:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error while trying to ping database:", err)
	}

	fmt.Println("Successfully connected to database")

	result, err := addRow(db, "PSP 3008", 10000)
	if err != nil {
		log.Fatal("Error in function \"addRow\":", err)
	}
	fmt.Println("Item with id =", result, "added successfully")

	result, err = addRow(db, "Electric guitar", 25000)
	if err != nil {
		log.Fatal("Error in function \"addRow\":", err)
	}
	fmt.Println("Item with id =", result, "added successfully")

	fmt.Println("All items in database:")

	for _, item := range selectAllRows(db) {
		fmt.Println("\t", item)
	}
}
