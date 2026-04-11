package handlers

import (
	taskdomain "example.com/taskservice/internal/domain/task"
)

func NewTaskDTO(task *taskdomain.Task) TaskDTO {
	return TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
