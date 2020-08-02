package initialozeapp

import (
	"fmt"

	"github.com/just-arun/server/internal/database"
)

// Init initialize seedin
func Init() {
	db, err := database.OpenDB()
	if err != nil {
		fmt.Println("Error %s", err)
	}
}
