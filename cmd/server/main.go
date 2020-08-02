package main

import (
	"fmt"
	"strconv"

	"github.com/just-arun/server/internal/database"
	"github.com/just-arun/server/internal/enums/userstatus"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var name string
	fmt.Println("enter your name")
	fmt.Scanln(&name)
	db, _ := database.OpenDB()
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
	statement.Exec()
	statement, _ = db.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	statement.Exec("Nic", name)
	rows, _ := db.Query("SELECT id, firstname, lastname FROM people")
	var id int
	var firstname string
	var lastname string
	for rows.Next() {
		rows.Scan(&id, &firstname, &lastname)
		fmt.Println(strconv.Itoa(id) + ": " + firstname + " " + lastname)
	}

	fmt.Println(userstatus.Active)
}
