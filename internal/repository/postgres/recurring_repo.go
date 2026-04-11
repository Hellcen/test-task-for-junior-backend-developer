package postgres

import "database/sql"

type RecurringRuleRepository struct {
	db *sql.DB
}

func NewRecurringRuleRepository(db *sql.DB) *RecurringRuleRepository {
	return &RecurringRuleRepository{
		db: db,
	}
}

func (r *RecurringRuleRepository) Create()
