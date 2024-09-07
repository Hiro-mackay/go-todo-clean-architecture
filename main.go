package main

import (
	"go-react-todo/controller"
	"go-react-todo/db"
	"go-react-todo/repository"
	"go-react-todo/router"
	"go-react-todo/usecase"
)

func main() {
	db := db.NewDB()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userController := controller.NewUserController(userUsecase)
	e := router.NewRouter(userController)
	e.Logger.Fatal(e.Start(":8080"))
}
