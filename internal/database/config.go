package database

import (
	"database/sql"
)

/*
OpenDB set up a connection to the database
*/
func OpenDB() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "./blog.db");
	return database, err;
}
