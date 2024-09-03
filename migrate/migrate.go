package main

import (
	"fmt"
	"go-react-todo/db"
	"go-react-todo/models"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&models.User{}, &models.Task{})
}
