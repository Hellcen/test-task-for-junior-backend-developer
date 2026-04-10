package recurring

import (
	"context"
	"example.com/taskservice/internal/domain/recurringrule"
	"example.com/taskservice/internal/transport/http/handlers"
	"time"
)

type Repository interface {
}

type Usecase interface {
	Create(ctx context.Context, input handlers.CreateInput) (*recurringrule.RecurringRule, error)
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

func (s *Service) Create(ctx context.Context, input handlers.CreateInput) (*recurringrule.RecurringRule, error) {
	return
}
