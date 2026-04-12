package handlers

import (
	"example.com/taskservice/internal/domain/recurringrule"
	recurringusecase "example.com/taskservice/internal/usecase/recurring"
)

func toCreateInputDTO(req RecurringRuleCreateRequest) recurringusecase.CreateInputDTO {
	return recurringusecase.CreateInputDTO{
		Title:          req.Title,
		Description:    req.Description,
		Status:         req.Status,
		RecurrenceType: req.RecurrenceType,
		Config:         req.Config,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
	}
}

func toUpdateInputDTO(req RecurringRuleUpdateRequest) recurringusecase.UpdateInputDTO {
	return recurringusecase.UpdateInputDTO{
		Title:          req.Title,
		Description:    req.Description,
		Status:         req.Status,
		RecurrenceType: req.RecurrenceType,
		Config:         req.Config,
		StartDate:      req.StartDate,
		EndDate:        req.EndDate,
	}
}

func toRecurringRuleResponse(rule *recurringrule.RecurringRule) RecurringRuleResponse {
	return RecurringRuleResponse{
		ID:             rule.ID,
		Title:          rule.Title,
		Description:    rule.Description,
		Status:         rule.Status,
		RecurrenceType: rule.RecurrenceType,
		Config:         rule.RecurrenceConfig,
		StartDate:      rule.StartDate,
		EndDate:        rule.EndDate,
		CreatedAt:      rule.CreatedAt,
		UpdatedAt:      rule.UpdatedAt,
	}
}
