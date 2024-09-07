package repository

import (
	"fmt"
	"go-react-todo/server/models"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ITaskRepository interface {
	GetAllTasks(tasks *[]models.Task, userId uint) error
	GetTaskById(task *models.Task, userId uint, taskId uint) error
	CreateTask(task *models.Task) error
	UpdateTask(task *models.Task, userId uint) error
	DeleteTask(userId uint, taskId uint) error
}

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{db}
}

func (r *TaskRepository) GetAllTasks(tasks *[]models.Task, userId uint) error {
	if err := r.db.Where("user_id = ?", userId).Find(tasks).Error; err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) GetTaskById(task *models.Task, userId uint, taskId uint) error {
	if err := r.db.Where("user_id = ? AND id = ?", userId, taskId).First(task).Error; err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	if err := r.db.Create(task).Error; err != nil {
		return err
	}

	return nil
}

func (r *TaskRepository) UpdateTask(task *models.Task, userId uint) error {
	result := r.db.Model(&task).Clauses(clause.Returning{}).Where("user_id = ? AND id = ?", userId, task.ID).Update("title", task.Title)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}

func (r *TaskRepository) DeleteTask(userId uint, taskId uint) error {
	result := r.db.Where("user_id = ? AND id = ?", userId, taskId).Delete(&models.Task{})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected < 1 {
		return fmt.Errorf("object does not exist")
	}

	return nil
}
