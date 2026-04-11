package postgres

import (
	"context"
	"encoding/json"
	"example.com/taskservice/internal/domain/recurringrule"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type RecurringRuleRepository struct {
	pool *pgxpool.Pool
}

func NewRecurringRuleRepository(db *pgxpool.Pool) *RecurringRuleRepository {
	return &RecurringRuleRepository{
		pool: db,
	}
}

func (r *RecurringRuleRepository) Create(ctx context.Context, rule *recurringrule.RecurringRule) (*recurringrule.RecurringRule, error) {
	query := `
		INSERT INTO recurring_rules (
			title, description, status, recurrence_type,
			interval_days, day_of_month, parity, specific_dates,
			start_date, end_date, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
		RETURNING id
	`

	var specificDate []byte
	if rule.RecurrenceSpecificDates != nil {
		specificDate, _ = json.Marshal(rule.RecurrenceSpecificDates)
	}

	err := r.pool.QueryRow(ctx, query,
		rule.Title, rule.Description, rule.Status, rule.RecurrenceType,
		rule.RecurrenceInterval, rule.RecurrenceDayOfMonth, rule.RecurrenceParity, specificDate,
		rule.RecurrenceStartDate, rule.RecurrenceEndDate, rule.CreatedAt, rule.UpdatedAt,
	).Scan(&rule.Id)

	if err != nil {
		return nil, fmt.Errorf("create recurring rule: %w", err)
	}

	return rule, nil
}

func (r *RecurringRuleRepository) List(ctx context.Context) ([]recurringrule.RecurringRule, error) {
	query := `
		SELECT id, title, description, status, recurrence_type,
		       interval_days, day_of_month, parity, specific_dates,
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
		var specificDatesJSON []byte

		err := rows.Scan(
			&rule.Id, &rule.Title, &rule.Description, &rule.Status, &rule.RecurrenceType,
			&rule.RecurrenceInterval, &rule.RecurrenceDayOfMonth, &rule.RecurrenceParity, &specificDatesJSON,
			&rule.RecurrenceStartDate, &rule.RecurrenceEndDate, &rule.CreatedAt, &rule.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan recurring rule: %w", err)
		}

		if specificDatesJSON != nil {
			json.Unmarshal(specificDatesJSON, &rule.RecurrenceSpecificDates)
		}

		rules = append(rules, rule)
	}

	return rules, nil
}
