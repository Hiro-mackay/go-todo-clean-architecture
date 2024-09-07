package usecase

import (
	"go-react-todo/models"
	"go-react-todo/repository"
)

type ITaskUsecase interface {
	GetAllTasks(userId uint) ([]models.TaskResponse, error)
	GetTaskById(userId uint, taskId uint) (models.TaskResponse, error)
	CreateTask(task models.Task) (models.TaskResponse, error)
	UpdateTask(userId uint, task models.Task) (models.TaskResponse, error)
	DeleteTask(userId uint, taskId uint) error
}

type TaskUsecase struct {
	tr repository.ITaskRepository
}

func NewTaskUsecase(tr repository.ITaskRepository) ITaskUsecase {
	return &TaskUsecase{tr}
}

func (u *TaskUsecase) GetAllTasks(userId uint) ([]models.TaskResponse, error) {
	var tasks []models.Task
	err := u.tr.GetAllTasks(&tasks, userId)
	if err != nil {
		return nil, err
	}

	var taskResponses []models.TaskResponse
	for _, task := range tasks {
		taskResponses = append(taskResponses, models.TaskResponse{
			ID:        task.ID,
			Title:     task.Title,
			CreatedAt: task.CreatedAt,
			UpdatedAt: task.UpdatedAt,
		})
	}

	return taskResponses, nil
}

func (u *TaskUsecase) GetTaskById(userId uint, taskId uint) (models.TaskResponse, error) {
	var task models.Task
	err := u.tr.GetTaskById(&task, userId, taskId)
	if err != nil {
		return models.TaskResponse{}, err
	}

	taskResponse := models.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskResponse, nil
}

func (u *TaskUsecase) CreateTask(task models.Task) (models.TaskResponse, error) {
	err := u.tr.CreateTask(&task)
	if err != nil {
		return models.TaskResponse{}, err
	}

	taskResponse := models.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskResponse, nil
}

func (u *TaskUsecase) UpdateTask(userId uint, task models.Task) (models.TaskResponse, error) {
	err := u.tr.UpdateTask(&task, userId)
	if err != nil {
		return models.TaskResponse{}, err
	}

	taskResponse := models.TaskResponse{
		ID:        task.ID,
		Title:     task.Title,
		CreatedAt: task.CreatedAt,
		UpdatedAt: task.UpdatedAt,
	}

	return taskResponse, nil
}

func (u *TaskUsecase) DeleteTask(userId uint, taskId uint) error {
	err := u.tr.DeleteTask(userId, taskId)
	if err != nil {
		return err
	}

	return nil
}
