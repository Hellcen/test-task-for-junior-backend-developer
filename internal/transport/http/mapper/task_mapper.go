package mapper

import (
	taskdomain "example.com/taskservice/internal/domain/task"
	"example.com/taskservice/internal/transport/http/handlers"
)

func NewTaskDTO(task *taskdomain.Task) handlers.TaskDTO {
	return handlers.TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
