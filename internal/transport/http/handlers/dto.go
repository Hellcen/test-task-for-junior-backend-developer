package handlers

import (
	"encoding/json"
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

type RecurringRuleCreateRequest struct {
	Title          string                       `json:"title"`
	Description    string                       `json:"description"`
	Status         recurringrule.Status         `json:"status,omitempty"`
	RecurrenceType recurringrule.RecurrenceType `json:"recurrence_type"`
	Config         json.RawMessage              `json:"config"`
	StartDate      time.Time                    `json:"start_date"`
	EndDate        *time.Time                   `json:"end_date,omitempty"`
}

type RecurringRuleUpdateRequest struct {
	Title          string                       `json:"title,omitempty"`
	Description    string                       `json:"description,omitempty"`
	Status         recurringrule.Status         `json:"status,omitempty"`
	RecurrenceType recurringrule.RecurrenceType `json:"recurrence_type,omitempty"`
	Config         json.RawMessage              `json:"config,omitempty"`
	StartDate      time.Time                    `json:"start_date,omitempty"`
	EndDate        *time.Time                   `json:"end_date,omitempty"`
}

type RecurringRuleResponse struct {
	ID             int64                        `json:"id"`
	Title          string                       `json:"title"`
	Description    string                       `json:"description"`
	Status         recurringrule.Status         `json:"status"`
	RecurrenceType recurringrule.RecurrenceType `json:"recurrence_type"`
	Config         json.RawMessage              `json:"config"`
	StartDate      time.Time                    `json:"start_date"`
	EndDate        *time.Time                   `json:"end_date,omitempty"`
	CreatedAt      time.Time                    `json:"created_at"`
	UpdatedAt      time.Time                    `json:"updated_at"`
}
