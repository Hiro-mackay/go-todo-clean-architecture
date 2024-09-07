package controller

import (
	"go-react-todo/server/models"
	"go-react-todo/server/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type ITaskController interface {
	GetAllTasks(c echo.Context) error
	GetTaskById(c echo.Context) error
	CreateTask(c echo.Context) error
	UpdateTask(c echo.Context) error
	DeleteTask(c echo.Context) error
}

type TaskController struct {
	tu usecase.ITaskUsecase
}

func NewTaskController(tc usecase.ITaskUsecase) ITaskController {
	return &TaskController{tc}
}

func (c *TaskController) GetAllTasks(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId, ok := claims["user_id"]
	if !ok {
		// userId が存在しない場合のエラーハンドリング
		log.Println("userId not found in claims")
		return ctx.JSON(http.StatusInternalServerError, "userId not found in claims")
	}

	taskResponses, err := c.tu.GetAllTasks(uint(userId.(float64)))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, taskResponses)
}

func (c *TaskController) GetTaskById(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskId, _ := strconv.Atoi(ctx.Param("id"))
	taskResponse, err := c.tu.GetTaskById(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, taskResponse)
}

func (c *TaskController) CreateTask(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var task models.Task
	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	task.UserId = uint(userId.(float64))
	taskResponse, err := c.tu.CreateTask(task)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusCreated, taskResponse)
}

func (c *TaskController) UpdateTask(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var task models.Task
	if err := ctx.Bind(&task); err != nil {
		return ctx.JSON(http.StatusBadRequest, err.Error())
	}

	taskResponse, err := c.tu.UpdateTask(uint(userId.(float64)), task)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, taskResponse)
}

func (c *TaskController) DeleteTask(ctx echo.Context) error {
	user := ctx.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	taskId, _ := strconv.Atoi(ctx.Param("id"))

	err := c.tu.DeleteTask(uint(userId.(float64)), uint(taskId))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	return ctx.NoContent(http.StatusOK)
}
