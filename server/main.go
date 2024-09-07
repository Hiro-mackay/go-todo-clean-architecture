package main

import (
	"go-react-todo/server/controller"
	"go-react-todo/server/db"
	"go-react-todo/server/repository"
	"go-react-todo/server/router"
	"go-react-todo/server/usecase"
	"go-react-todo/server/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

	taskValidator := validator.NewTaskValidator()
	taskRepository := repository.NewTaskRepository(db)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
