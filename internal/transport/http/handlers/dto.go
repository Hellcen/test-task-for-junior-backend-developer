package handlers

import (
	"example.com/taskservice/internal/domain/recurringrule"
	"time"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type taskMutationDTO struct {
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
}

type taskDTO struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

func newTaskDTO(task *taskdomain.Task) taskDTO {
	return taskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

// CreateInput DTO для создания
type CreateInput struct {
	Title          string
	Description    string
	Status         recurringrule.Status
	RecurrenceType recurringrule.RecurrenceType
	IntervalDays   *int
	DayOfMonth     *int
	Parity         *recurringrule.Parity
	SpecificDates  recurringrule.JSONDates
	StartDate      time.Time
	EndDate        *time.Time
}
