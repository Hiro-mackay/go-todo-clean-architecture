package main

import (
	"fmt"
	"go-react-todo/server/db"
	"go-react-todo/server/models"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&models.User{}, &models.Task{})
}
