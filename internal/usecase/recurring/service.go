package recurring

import (
	"context"
	"errors"
	"fmt"
	"time"

	"example.com/taskservice/internal/domain/recurringrule"
)

var (
	ErrInvalidInput = errors.New("invalid input")
)

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

func (s *Service) Create(ctx context.Context, input CreateInputDTO) (*recurringrule.RecurringRule, error) {
	validated, err := ValidateCreateInput(input)
	if err != nil {
		return nil, err
	}

	now := s.now()

	rule := &recurringrule.RecurringRule{
		Title:            validated.Title,
		Description:      validated.Description,
		Status:           validated.Status,
		RecurrenceType:   validated.RecurrenceType,
		RecurrenceConfig: validated.Config,
		StartDate:        validated.StartDate,
		EndDate:          validated.EndDate,
		CreatedAt:        now,
		UpdatedAt:        now,
	}

	created, err := s.repo.Create(ctx, rule)
	if err != nil {
		return nil, fmt.Errorf("create rule: %w", err)
	}

	return created, nil
}

func (s *Service) GetByID(ctx context.Context, id int64) (*recurringrule.RecurringRule, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.GetByID(ctx, id)
}

func (s *Service) List(ctx context.Context) ([]recurringrule.RecurringRule, error) {
	return s.repo.List(ctx)
}

func (s *Service) Delete(ctx context.Context, id int64) error {
	if id <= 0 {
		return fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	return s.repo.Delete(ctx, id)
}

func (s *Service) Update(ctx context.Context, id int64, input UpdateInputDTO) (*recurringrule.RecurringRule, error) {
	if id <= 0 {
		return nil, fmt.Errorf("%w: id must be positive", ErrInvalidInput)
	}

	existing, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	validated, err := ValidateUpdateInput(input, *existing)
	if err != nil {
		return nil, err
	}

	existing.Title = validated.Title
	existing.Description = validated.Description
	existing.Status = validated.Status
	existing.RecurrenceType = validated.RecurrenceType
	existing.RecurrenceConfig = validated.Config
	existing.StartDate = validated.StartDate
	existing.EndDate = validated.EndDate
	existing.UpdatedAt = s.now()

	updated, err := s.repo.Update(ctx, existing)
	if err != nil {
		return nil, fmt.Errorf("update rule: %w", err)
	}

	return updated, nil
}

func (s *Service) UpdateStatus(ctx context.Context, id int64, status recurringrule.Status) error {

}
