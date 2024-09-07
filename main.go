package main

import (
	"go-react-todo/controller"
	"go-react-todo/db"
	"go-react-todo/repository"
	"go-react-todo/router"
	"go-react-todo/usecase"
	"go-react-todo/validator"
)

func main() {
	db := db.NewDB()

	userValidatotr := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidatotr)
	userController := controller.NewUserController(userUsecase)

	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
