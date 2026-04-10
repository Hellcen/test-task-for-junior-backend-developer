CREATE TABLE IF NOT EXISTS recurring_rules
(
    id              BIGSERIAL PRIMARY KEY,
    title           TEXT      NOT NULL,
    description     TEXT      NOT NULL,
    status          TEXT      NOT NULL CHECK (status IN ('active', 'paused', 'archived')),
    recurrence_type TEXT      NOT NULL CHECK (recurrence_type IN ('daily', 'monthly', 'evenodd', 'specific')),

    -- Daily
    interval_days   INT CHECK (interval_days > 0),

    -- Monthly
    day_of_month    INT CHECK (day_of_month BETWEEN 1 AND 30),

    -- EvenOdd
    parity          TEXT CHECK (parity IN ('even', 'odd')),

    -- Specific
    specific_dates  JSONB,

    -- Общие поля
    start_date      DATE      NOT NULL,
    end_date        DATE,
    created_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_recurring_rules_status ON recurring_rules (status);
CREATE INDEX idx_recurring_rules_type ON recurring_rules (recurrence_type);

ALTER TABLE tasks
    ADD COLUMN IF NOT EXISTS recurring_rule_id BIGINT REFERENCES recurring_rules (id) ON DELETE SET NULL,
    ADD COLUMN IF NOT EXISTS generation_date   DATE;

CREATE INDEX IF NOT EXISTS idx_tasks_recurring_rule_id ON tasks (recurring_rule_id);