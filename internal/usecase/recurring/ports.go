package recurring

import (
	"context"
	"encoding/json"
	"time"

	"example.com/taskservice/internal/domain/recurringrule"
)

type Repository interface {
	Create(ctx context.Context, rule *recurringrule.RecurringRule) (*recurringrule.RecurringRule, error)
	GetByID(ctx context.Context, id int64) (*recurringrule.RecurringRule, error)
	List(ctx context.Context) ([]recurringrule.RecurringRule, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, rule *recurringrule.RecurringRule) (*recurringrule.RecurringRule, error)
	UpdateStatus(ctx context.Context, id int64, status recurringrule.Status) error
}

type Usecase interface {
	Create(ctx context.Context, input CreateInputDTO) (*recurringrule.RecurringRule, error)
	GetByID(ctx context.Context, id int64) (*recurringrule.RecurringRule, error)
	List(ctx context.Context) ([]recurringrule.RecurringRule, error)
	Delete(ctx context.Context, id int64) error
	Update(ctx context.Context, id int64, input UpdateInputDTO) (*recurringrule.RecurringRule, error)
	UpdateStatus(ctx context.Context, id int64, status recurringrule.Status) error
}

type CreateInputDTO struct {
	Title          string
	Description    string
	Status         recurringrule.Status
	RecurrenceType recurringrule.RecurrenceType
	Config         json.RawMessage
	StartDate      time.Time
	EndDate        *time.Time
}

type UpdateInputDTO struct {
	Title          string
	Description    string
	Status         recurringrule.Status
	RecurrenceType recurringrule.RecurrenceType
	Config         json.RawMessage
	StartDate      time.Time
	EndDate        *time.Time
}
