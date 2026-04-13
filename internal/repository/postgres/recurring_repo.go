package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	"example.com/taskservice/internal/domain/recurringrule"
)

type RecurringRuleRepository struct {
	pool *pgxpool.Pool
}

func NewRecurringRuleRepository(pool *pgxpool.Pool) *RecurringRuleRepository {
	return &RecurringRuleRepository{pool: pool}
}

func (r *RecurringRuleRepository) Create(ctx context.Context, rule *recurringrule.RecurringRule) (*recurringrule.RecurringRule, error) {
	query := `
		INSERT INTO recurring_rules (
			title, description, status, recurrence_type, recurrence_config,
			start_date, end_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id
	`

	err := r.pool.QueryRow(ctx, query,
		rule.Title, rule.Description, rule.Status, rule.RecurrenceType,
		rule.RecurrenceConfig, rule.StartDate, rule.EndDate,
		rule.CreatedAt, rule.UpdatedAt,
	).Scan(&rule.ID)

	if err != nil {
		return nil, fmt.Errorf("create recurring rule: %w", err)
	}

	return rule, nil
}

func (r *RecurringRuleRepository) GetByID(ctx context.Context, id int64) (*recurringrule.RecurringRule, error) {
	query := `
		SELECT id, title, description, status, recurrence_type, recurrence_config,
		       start_date, end_date, created_at, updated_at
		FROM recurring_rules WHERE id = $1
	`

	rule := &recurringrule.RecurringRule{}

	err := r.pool.QueryRow(ctx, query, id).Scan(
		&rule.ID, &rule.Title, &rule.Description, &rule.Status, &rule.RecurrenceType,
		&rule.RecurrenceConfig, &rule.StartDate, &rule.EndDate,
		&rule.CreatedAt, &rule.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, recurringrule.ErrNotFound
		}
		return nil, fmt.Errorf("get recurring rule by id: %w", err)
	}

	return rule, nil
}

func (r *RecurringRuleRepository) List(ctx context.Context) ([]recurringrule.RecurringRule, error) {
	query := `
		SELECT id, title, description, status, recurrence_type, recurrence_config,
		       start_date, end_date, created_at, updated_at
		FROM recurring_rules ORDER BY created_at DESC
	`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list recurring rules: %w", err)
	}
	defer rows.Close()

	var rules []recurringrule.RecurringRule

	for rows.Next() {
		var rule recurringrule.RecurringRule

		err := rows.Scan(
			&rule.ID, &rule.Title, &rule.Description, &rule.Status, &rule.RecurrenceType,
			&rule.RecurrenceConfig, &rule.StartDate, &rule.EndDate,
			&rule.CreatedAt, &rule.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan recurring rule: %w", err)
		}

		rules = append(rules, rule)
	}

	return rules, nil
}

func (r *RecurringRuleRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM recurring_rules WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete recurring rule: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return recurringrule.ErrNotFound
	}

	return nil
}

func (r *RecurringRuleRepository) Update(ctx context.Context, rule *recurringrule.RecurringRule) (*recurringrule.RecurringRule, error) {
	query := `
		UPDATE recurring_rules SET
			title = $1, description = $2, status = $3, recurrence_type = $4,
			recurrence_config = $5, start_date = $6, end_date = $7, updated_at = $8
		WHERE id = $9
	`

	cmdTag, err := r.pool.Exec(ctx, query,
		rule.Title, rule.Description, rule.Status, rule.RecurrenceType,
		rule.RecurrenceConfig, rule.StartDate, rule.EndDate, rule.UpdatedAt, rule.ID,
	)

	if err != nil {
		return nil, fmt.Errorf("update recurring rule: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, recurringrule.ErrNotFound
	}

	return rule, nil
}

func (r *RecurringRuleRepository) UpdateStatus(ctx context.Context, id int64, status recurringrule.Status) error {
	query := `UPDATE recurring_rules SET status = $1, updated_at = $2 WHERE id = $3`

	cmdTag, err := r.pool.Exec(ctx, query, status, time.Now(), id)
	if err != nil {
		return fmt.Errorf("update recurring rule: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return recurringrule.ErrNotFound
	}

	return nil
}
