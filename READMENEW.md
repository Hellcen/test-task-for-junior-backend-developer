# Task Service

Сервис для управления задачами с поддержкой периодических (рекуррентных) задач. Реализован на Go с использованием чистой
архитектуры, паттернов проектирования и современных подходов к конкурентности.

## Содержание

- [Особенности](#особенности)
- [Паттерны проектирования](#паттерны-проектирования)
- [Паттерны параллельности](#паттерны-параллельности)
- [База данных](#база-данных)
- [API Endpoints](#api-endpoints)
- [Мониторинг](#мониторинг)
- [Быстрый запуск](#быстрый-запуск)
- [Примеры запросов](#примеры-запросов)

## Особенности

- **CRUD операции** для обычных задач
- **Периодические задачи** с поддержкой 4 типов:
    - Ежедневные (с настраиваемым интервалом)
    - Ежемесячные (определённое число месяца)
    - Чётные/нечётные дни месяца
    - Конкретные даты
- **Чистая архитектура** (Clean Architecture)
- **Паттерны проектирования**: Strategy, Repository, DTO, Factory
- **Паттерны параллельности**: Worker Pool, Producer-Consumer
- **Метрики и мониторинг**: Prometheus + Grafana
- **JSONB в PostgreSQL** для гибкого хранения конфигураций
- **Docker Compose** для простого развёртывания

# Паттерны проектирования

### Strategy Pattern (Стратегия)

Используется для генерации дат для разных типов периодичности.

```go
type DateStrategy interface {
GenerateDates(rule RecurringRule) ([]time.Time, error)
}

type DailyConfig struct {
    IntervalDays int `json:"interval_days"`
}

type MonthlyConfig struct {
    DayOfMonth int `json:"day_of_month"`
}

type EvenOddConfig struct {
    Parity string `json:"parity"` // "even" или "odd"
}

type SpecificConfig struct {
    Dates []time.Time `json:"dates"`
}
```

### Repository Pattern (Репозиторий)

Изолирует бизнес-логику от конкретной реализации базы данных.

```go
type Repository interface {
Create(ctx context.Context, rule *RecurringRule) (*RecurringRule, error)
GetByID(ctx context.Context, id int64) (*RecurringRule, error)
Update(ctx context.Context, rule *RecurringRule) (*RecurringRule, error)
Delete(ctx context.Context, id int64) error
List(ctx context.Context) ([]RecurringRule, error)
}
```

### DTO Pattern (Data Transfer Object)

Разные DTO для разных слоёв:

| Слой    | DTO                        | Назначение               |
|:--------|:---------------------------|:-------------------------|		
| HTTP	   | RecurringRuleCreateRequest | Парсинг JSON             |
| UseCase | CreateInputDTO	            | Передача в бизнес-логику |
| Domain  | 	RecurringRule             | Хранение в БД            |

### Factory Pattern (Фабрика)

Создание нужной стратегии на основе типа периодичности.

```go
func (f *StrategyFactory) GetStrategy(rule RecurringRule) (DateStrategy, error) {
switch rule.RecurrenceType {
case TypeDaily:    return &DailyConfig{}, nil
case TypeMonthly:  return &MonthlyConfig{}, nil
case TypeEvenOdd:  return &EvenOddConfig{}, nil
case TypeSpecific: return &SpecificConfig{}, nil
}
}
```

## Паттерны параллельности

### Worker Pool Pattern

Ограниченное количество воркеров обрабатывает задачи из общего канала.

```go
func (g *TaskGenerator) GenerateTasksParallel(rules []RecurringRule) {
jobs := make(chan RecurringRule, len(rules))
var wg sync.WaitGroup


for i := 0; i < workerCount; i++ {
wg.Add(1)
go g.worker(jobs, &results, &mu, &wg)
}
//...
}
```

Зачем: Контроль использования ресурсов, предотвращение создания тысяч горутин.

### Background Worker Pattern

Фоновый сбор метрик в отдельной горутине.

```go
go metrics.CollectMetrics(recurringRuleUsecase)
```

## База данных

### Почему JSONB?

Проблема: нужно хранить конфигурации 4 разных типов периодичности.

| Вариант           | Плюсы           | Минусы              |
|:------------------|:----------------|:--------------------|
| Отдельные поля    | Простые запросы | Много Null-значений |
| Отдельные таблицы | Нормализация    | Много Join операций |
| JSONB             | Нет null        | Валидация в коде    |

### Примеры конфигураций

```JSON
// daily
{"interval_days": 2}

// monthly
{"day_of_month": 5}

// evenodd
{"parity": "even"}

// specific
{"dates": ["2026-04-15", "2026-05-01"]}
```

### Индексы

```sql
    CREATE INDEX idx_recurring_rules_status ON recurring_rules(status);
    CREATE INDEX idx_recurring_rules_type ON recurring_rules(recurrence_type);
    CREATE INDEX idx_recurring_rules_config ON recurring_rules USING GIN (recurrence_config);
```

## API Endpoints

### Tasks (обычные задачи)

| Метод | URL           | Описание |
|:------|:--------------|:---------|
| POST  | /api/v1/tasks | Создать задачу         |
| GET   | /api/v1/tasks             |Получить все задачи|
| GET   |/api/v1/tasks/{id}|Получить задачу по ID|
| PUT   |/api/v1/tasks/{id}|Обновить задачу|
| DELETE     |/api/v1/tasks/{id}|Удалить задачу|

### Recurring Rules (периодические правила)

| Метод  | URL                         | Описание |
|:-------|:----------------------------|:---------|
| POST   | /api/v1/recurringrule-rules | Создать правило         |
| GET    | /api/v1/recurringrule-rules                            |Получить все правила|
| GET    |/api/v1/recurringrule-rules/{id}|Получить правило по ID|
| PUT    |/api/v1/recurringrule-rules/{id}|Обновить правило|
| DELETE |/api/v1/recurringrule-rules/{id}|	Удалить правило|
| PATCH  |/api/v1/recurringrule-rules/{id}/status|Обновить статус|

### Системные

| Метод | URL | Описание |
|:------|:----|:---------|
| GET   |/health|Health check|
| GET   |/metrics|Prometheus метрики|
| GET   |/swagger/|Swagger UI|
| GET   |/swagger/openapi.json|OpenAPI спецификация|

## Мониторинг 

### Метрики

| Метрика                         | Тип       | Описание                 |
|:--------------------------------|:----------|:-------------------------|
| `http_requests_total`           | Counter   | Количество HTTP запросов |
| `http_request_duration_seconds` | Histogram | Длительность запросов    |
| `active_recurring_rules_total`  | Gauge     | Активные правила         |
| `total_recurring_rules_total`   | Gauge     | Всего правил             |
| `recurring_rules_by_type_total` | Gauge     | Распределение по типам   |

### Быстрый запуск

```bash
 git clone <repository-url>
 
 docker-compose up --build
 
 docker-compose down -v
```

## Примеры запросов

### Создание правила (daily)

```bash
curl -X POST http://localhost:8080/api/v1/recurringrule-rules \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Ежедневная проверка",
    "description": "Проверять систему каждый день",
    "recurrence_type": "daily",
    "config": {"interval_days": 2},
    "start_date": "2026-04-13T00:00:00Z"
  }'
```

### Создание правила (monthly)

```bash
curl -X POST http://localhost:8080/api/v1/recurringrule-rules \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Формирование отчёта",
    "recurrence_type": "monthly",
    "config": {"day_of_month": 5},
    "start_date": "2026-04-13T00:00:00Z"
  }'
```

### Созлание правила (evenodd)

```bash
curl -X POST http://localhost:8080/api/v1/recurringrule-rules \
-H "Content-Type: application/json" \ 
-d '{
  "title": "Инвентаризация", 
  "description": "По чётным дням", 
  "recurrence_type": "evenodd", 
  "config": {"parity": "even"}, 
  "start_date": "2026-04-13T00:00:00Z"
}'
```

### Создание правила (specific)

```bash
curl -X POST http://localhost:8080/api/v1/recurringrule-rules \
 -H "Content-Type: application/json" \
 -d '{
    "title": "Особые даты", 
    "description": "Только в указанные даты", 
    "recurrence_type": "specific", 
    "config": {"dates": ["2026-04-15T00:00:00Z", "2026-05-01T00:00:00Z"]}, 
    "start_date": "2026-04-13T00:00:00Z"
  }'
```

### Получить все правила

```bash
curl -X GET http://localhost:8080/api/v1/recurringrule-rules
```

### Обновить статус правила

```bash
curl -X PATCH http://localhost:8080/api/v1/recurringrule-rules/1/status \
  -H "Content-Type: application/json" \
  -d '{"status": "paused"}'
```

### Создать обычную задачу

```bash
curl -X POST http://localhost:8080/api/v1/tasks \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Подготовить релиз",
    "description": "Собрать release notes",
    "status": "new"
  }'
```