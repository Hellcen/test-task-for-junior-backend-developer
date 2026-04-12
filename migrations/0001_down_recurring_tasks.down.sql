ALTER TABLE tasks
    DROP COLUMN IF EXISTS recurring_rule_id,
    DROP COLUMN IF EXISTS generation_date;

DROP INDEX IF EXISTS idx_tasks_recurring_rule_id;