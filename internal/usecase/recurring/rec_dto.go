package recurring

import (
	"time"

	"example.com/taskservice/internal/domain/recurringrule"
)

// CreateRecurringRuleInputDTO для создания правила
type CreateRecurringRuleInputDTO struct {
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
}

// UpdateRecurringRuleInputDTO для обновления правила
type UpdateRecurringRuleInputDTO struct {
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
}

// RecurringRuleResponseDTO для ответа (если нужен общий DTO для usecase)
type RecurringRuleDTO struct {
	ID                      int64
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
