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

// RecurringRuleCreateRequest для создания правила (HTTP request)
type RecurringRuleCreateRequest struct {
	Title                   string                       `json:"title"`
	Description             string                       `json:"description"`
	Status                  recurringrule.Status         `json:"status,omitempty"`
	RecurrenceType          recurringrule.RecurrenceType `json:"recurrence_type"`
	RecurrenceInterval      *int                         `json:"interval_days,omitempty"`
	RecurrenceDayOfMonth    *int                         `json:"day_of_month,omitempty"`
	RecurrenceParity        *recurringrule.Parity        `json:"parity,omitempty"`
	RecurrenceSpecificDates []time.Time                  `json:"specific_dates,omitempty"`
	RecurrenceStartDate     time.Time                    `json:"start_date"`
	RecurrenceEndDate       *time.Time                   `json:"end_date,omitempty"`
}

// RecurringRuleUpdateRequest для обновления правила (HTTP request)
type RecurringRuleUpdateRequest struct {
	Title                   string                       `json:"title,omitempty"`
	Description             string                       `json:"description,omitempty"`
	Status                  recurringrule.Status         `json:"status,omitempty"`
	RecurrenceType          recurringrule.RecurrenceType `json:"recurrence_type,omitempty"`
	RecurrenceInterval      *int                         `json:"interval_days,omitempty"`
	RecurrenceDayOfMonth    *int                         `json:"day_of_month,omitempty"`
	RecurrenceParity        *recurringrule.Parity        `json:"parity,omitempty"`
	RecurrenceSpecificDates []time.Time                  `json:"specific_dates,omitempty"`
	RecurrenceStartDate     time.Time                    `json:"start_date,omitempty"`
	RecurrenceEndDate       *time.Time                   `json:"end_date,omitempty"`
}

// RecurringRuleResponse для ответа (HTTP response)
type RecurringRuleResponse struct {
	ID                      int64                        `json:"id"`
	Title                   string                       `json:"title"`
	Description             string                       `json:"description"`
	Status                  recurringrule.Status         `json:"status"`
	RecurrenceType          recurringrule.RecurrenceType `json:"recurrence_type"`
	RecurrenceInterval      *int                         `json:"interval_days,omitempty"`
	RecurrenceDayOfMonth    *int                         `json:"day_of_month,omitempty"`
	RecurrenceParity        *recurringrule.Parity        `json:"parity,omitempty"`
	RecurrenceSpecificDates []time.Time                  `json:"specific_dates,omitempty"`
	RecurrenceStartDate     time.Time                    `json:"start_date"`
	RecurrenceEndDate       *time.Time                   `json:"end_date,omitempty"`
	CreatedAt               time.Time                    `json:"created_at"`
	UpdatedAt               time.Time                    `json:"updated_at"`
}
