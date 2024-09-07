package validator

import (
	"go-react-todo/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type ITaskValidator interface {
	ValidateTask(task *models.Task) error
}

type TaskValidator struct{}

func NewTaskValidator() ITaskValidator {
	return &TaskValidator{}
}

func (v *TaskValidator) ValidateTask(task *models.Task) error {
	return validation.ValidateStruct(task,
		validation.Field(&task.Title, validation.Required, validation.RuneLength(1, 100).Error("Title must be between 1 and 100 characters")),
	)
}
