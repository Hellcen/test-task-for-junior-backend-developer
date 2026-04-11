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

type TaskDTO struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Status      taskdomain.Status `json:"status"`
	CreatedAt   time.Time         `json:"created_at"`
	UpdatedAt   time.Time         `json:"updated_at"`
}

// CreateInputDTO DTO для создания
type CreateInputDTO struct {
	Title                   string
	Description             string
	Status                  recurringrule.Status
	RecurrenceType          recurringrule.RecurrenceType
	RecurrenceInterval      *int
	RecurrenceDayOfMonth    *int
	RecurrenceParity        *recurringrule.Parity
	RecurrenceSpecificDates recurringrule.JSONDates
	RecurrenceStartDate     time.Time
	RecurrenceEndDate       *time.Time
	CreatedAt               time.Time
	UpdatedAt               time.Time
}
