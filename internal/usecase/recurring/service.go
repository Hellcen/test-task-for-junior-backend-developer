package recurring

import (
	"context"
	"example.com/taskservice/internal/domain/recurringrule"
	taskdomain "example.com/taskservice/internal/domain/task"
	"example.com/taskservice/internal/transport/http/handlers"
	"example.com/taskservice/internal/transport/http/mapper"
	"time"
)

type Repository interface {
	Create(context.Context, recurringrule.RecurringRule) error
}

type Usecase interface {
	Create(ctx context.Context, input handlers.CreateInputDTO) (*recurringrule.RecurringRule, error)
}

type Service struct {
	repo Repository
	now  func() time.Time
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
		now:  func() time.Time { return time.Now().UTC() },
	}
}

func (s *Service) Create(ctx context.Context, input handlers.CreateInputDTO) (*recurringrule.RecurringRule, error) {
	data, err := mapper.ValidateCreateInputDTO(input)

	if err != nil {
		return nil, err
	}

	rule := &recurringrule.RecurringRule{
		Title:                   data.Title,
		Description:             data.Description,
		Status:                  taskdomain.Status(data.Status),
		RecurrenceType:          data.RecurrenceType,
		RecurrenceInterval:      data.RecurrenceInterval,
		RecurrenceDayOfMonth:    data.RecurrenceDayOfMonth,
		RecurrenceParity:        data.RecurrenceParity,
		RecurrenceSpecificDates: data.RecurrenceSpecificDates,
		RecurrenceStartDate:     data.RecurrenceStartDate,
		RecurrenceEndDate:       data.RecurrenceEndDate,
		CreatedAt:               s.now(),
		UpdatedAt:               s.now(),
	}

	// TODO: repo.Create()...

	return rule, nil
}
