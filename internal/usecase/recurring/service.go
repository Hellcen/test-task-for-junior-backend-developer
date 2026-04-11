package recurring

import (
	"context"
	"errors"
	taskdomain "example.com/taskservice/internal/domain/task"
	"fmt"
	"strings"
	"time"

	"example.com/taskservice/internal/domain/recurringrule"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

type Repository interface {
	Create(ctx context.Context, rule *recurringrule.RecurringRule) (*recurringrule.RecurringRule, error)
	List(ctx context.Context) ([]recurringrule.RecurringRule, error)
}

type Usecase interface {
	Create(ctx context.Context, input CreateRecurringRuleInputDTO) (*recurringrule.RecurringRule, error)
	List(ctx context.Context) ([]recurringrule.RecurringRule, error)
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

func (s *Service) Create(ctx context.Context, input CreateRecurringRuleInputDTO) (*recurringrule.RecurringRule, error) {
	validated, err := ValidateCreateInput(input)
	if err != nil {
		return nil, err
	}

	now := s.now()

	rule := &recurringrule.RecurringRule{
		Title:                   validated.Title,
		Description:             validated.Description,
		Status:                  taskdomain.Status(validated.Status),
		RecurrenceType:          validated.RecurrenceType,
		RecurrenceInterval:      validated.RecurrenceInterval,
		RecurrenceDayOfMonth:    validated.RecurrenceDayOfMonth,
		RecurrenceParity:        validated.RecurrenceParity,
		RecurrenceSpecificDates: validated.RecurrenceSpecificDates,
		RecurrenceStartDate:     validated.RecurrenceStartDate,
		RecurrenceEndDate:       validated.RecurrenceEndDate,
		CreatedAt:               now,
		UpdatedAt:               now,
	}

	created, err := s.repo.Create(ctx, rule)
	if err != nil {
		return nil, fmt.Errorf("create rule: %w", err)
	}

	return created, nil
}

func (s *Service) List(ctx context.Context) ([]recurringrule.RecurringRule, error) {
	return s.repo.List(ctx)
}

// ValidateCreateInput валидация (остаётся здесь же)
func ValidateCreateInput(input CreateRecurringRuleInputDTO) (CreateRecurringRuleInputDTO, error) {
	input.Title = strings.TrimSpace(input.Title)
	input.Description = strings.TrimSpace(input.Description)

	if input.Title == "" {
		return CreateRecurringRuleInputDTO{}, fmt.Errorf("%w: title is required", ErrInvalidInput)
	}

	if input.Status == "" {
		input.Status = recurringrule.StatusActive
	}

	if !input.Status.Valid() {
		return CreateRecurringRuleInputDTO{}, fmt.Errorf("%w: invalid status", ErrInvalidInput)
	}

	if !input.RecurrenceType.Valid() {
		return CreateRecurringRuleInputDTO{}, fmt.Errorf("%w: invalid recurrence_type", ErrInvalidInput)
	}

	switch input.RecurrenceType {
	case recurringrule.TypeDaily:
		if input.RecurrenceInterval == nil || *input.RecurrenceInterval <= 0 {
			return CreateRecurringRuleInputDTO{}, fmt.Errorf("%w: interval_days must be > 0 for daily recurrence", ErrInvalidInput)
		}
	case recurringrule.TypeMonthly:
		if input.RecurrenceDayOfMonth == nil || *input.RecurrenceDayOfMonth < 1 || *input.RecurrenceDayOfMonth > 30 {
			return CreateRecurringRuleInputDTO{}, fmt.Errorf("%w: day_of_month must be between 1 and 30", ErrInvalidInput)
		}
	case recurringrule.TypeEvenOdd:
		if input.RecurrenceParity == nil || !input.RecurrenceParity.Valid() {
			return CreateRecurringRuleInputDTO{}, fmt.Errorf("%w: parity must be 'even' or 'odd'", ErrInvalidInput)
		}
	case recurringrule.TypeSpecific:
		if len(input.RecurrenceSpecificDates) == 0 {
			return CreateRecurringRuleInputDTO{}, fmt.Errorf("%w: specific_dates cannot be empty", ErrInvalidInput)
		}
	}

	if input.RecurrenceStartDate.IsZero() {
		input.RecurrenceStartDate = time.Now().UTC()
	}

	return input, nil
}
