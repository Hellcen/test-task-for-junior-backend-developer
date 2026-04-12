package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"

	taskdomain "example.com/taskservice/internal/domain/task"
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func New(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}

func (r *TaskRepository) Create(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error) {
	query := `
		INSERT INTO tasks (title, description, status, created_at, updated_at, recurring_rule_id, generation_date)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	err := r.pool.QueryRow(ctx, query,
		task.Title, task.Description, task.Status, task.CreatedAt, task.UpdatedAt,
		task.RecurringRuleID, task.GenerationDate,
	).Scan(&task.ID)

	if err != nil {
		return nil, fmt.Errorf("create task: %w", err)
	}

	return task, nil
}

func (r *TaskRepository) GetByID(ctx context.Context, id int64) (*taskdomain.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at, recurring_rule_id, generation_date 
	          FROM tasks WHERE id = $1`

	var task taskdomain.Task
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&task.ID, &task.Title, &task.Description, &task.Status,
		&task.CreatedAt, &task.UpdatedAt, &task.RecurringRuleID, &task.GenerationDate,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, taskdomain.ErrNotFound
		}
		return nil, fmt.Errorf("get task by id: %w", err)
	}

	return &task, nil
}

func (r *TaskRepository) Update(ctx context.Context, task *taskdomain.Task) (*taskdomain.Task, error) {
	query := `
		UPDATE tasks SET 
			title = $1, description = $2, status = $3, updated_at = $4,
			recurring_rule_id = $5, generation_date = $6
		WHERE id = $7
	`

	cmdTag, err := r.pool.Exec(ctx, query,
		task.Title, task.Description, task.Status, task.UpdatedAt,
		task.RecurringRuleID, task.GenerationDate, task.ID,
	)

	if err != nil {
		return nil, fmt.Errorf("update task: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return nil, taskdomain.ErrNotFound
	}

	return task, nil
}

func (r *TaskRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM tasks WHERE id = $1`

	cmdTag, err := r.pool.Exec(ctx, query, id)
	if err != nil {
		return fmt.Errorf("delete task: %w", err)
	}

	if cmdTag.RowsAffected() == 0 {
		return taskdomain.ErrNotFound
	}

	return nil
}

func (r *TaskRepository) List(ctx context.Context) ([]taskdomain.Task, error) {
	query := `SELECT id, title, description, status, created_at, updated_at, recurring_rule_id, generation_date 
	          FROM tasks ORDER BY created_at DESC`

	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("list tasks: %w", err)
	}
	defer rows.Close()

	var tasks []taskdomain.Task
	for rows.Next() {
		var task taskdomain.Task
		err := rows.Scan(
			&task.ID, &task.Title, &task.Description, &task.Status,
			&task.CreatedAt, &task.UpdatedAt, &task.RecurringRuleID, &task.GenerationDate,
		)
		if err != nil {
			return nil, fmt.Errorf("scan task: %w", err)
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
