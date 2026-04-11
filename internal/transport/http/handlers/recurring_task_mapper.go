package handlers

import (
	"example.com/taskservice/internal/domain/recurringrule"
	"example.com/taskservice/internal/usecase/recurring"
)

// toCreateInputDTO конвертирует HTTP request → Usecase DTO
func toCreateInputDTO(req RecurringRuleCreateRequest) recurring.CreateRecurringRuleInputDTO {
	return recurring.CreateRecurringRuleInputDTO{
		Title:                   req.Title,
		Description:             req.Description,
		Status:                  req.Status,
		RecurrenceType:          req.RecurrenceType,
		RecurrenceInterval:      req.RecurrenceInterval,
		RecurrenceDayOfMonth:    req.RecurrenceDayOfMonth,
		RecurrenceParity:        req.RecurrenceParity,
		RecurrenceSpecificDates: req.RecurrenceSpecificDates,
		RecurrenceStartDate:     req.RecurrenceStartDate,
		RecurrenceEndDate:       req.RecurrenceEndDate,
	}
}

// toUpdateInputDTO конвертирует HTTP request → Usecase DTO
func toUpdateInputDTO(req RecurringRuleUpdateRequest) recurring.UpdateRecurringRuleInputDTO {
	return recurring.UpdateRecurringRuleInputDTO{
		Title:                   req.Title,
		Description:             req.Description,
		Status:                  req.Status,
		RecurrenceType:          req.RecurrenceType,
		RecurrenceInterval:      req.RecurrenceInterval,
		RecurrenceDayOfMonth:    req.RecurrenceDayOfMonth,
		RecurrenceParity:        req.RecurrenceParity,
		RecurrenceSpecificDates: req.RecurrenceSpecificDates,
		RecurrenceStartDate:     req.RecurrenceStartDate,
		RecurrenceEndDate:       req.RecurrenceEndDate,
	}
}

// toRecurringRuleResponse конвертирует Domain → HTTP response
func toRecurringRuleResponse(rule *recurringrule.RecurringRule) RecurringRuleResponse {
	return RecurringRuleResponse{
		ID:                      int64(rule.Id),
		Title:                   rule.Title,
		Description:             rule.Description,
		Status:                  recurringrule.Status(rule.Status),
		RecurrenceType:          rule.RecurrenceType,
		RecurrenceInterval:      rule.RecurrenceInterval,
		RecurrenceDayOfMonth:    rule.RecurrenceDayOfMonth,
		RecurrenceParity:        rule.RecurrenceParity,
		RecurrenceSpecificDates: rule.RecurrenceSpecificDates,
		RecurrenceStartDate:     rule.RecurrenceStartDate,
		RecurrenceEndDate:       rule.RecurrenceEndDate,
		CreatedAt:               rule.CreatedAt,
		UpdatedAt:               rule.UpdatedAt,
	}
}

// toRecurringRuleResponseList конвертирует слайс Domain → слайс HTTP response
func toRecurringRuleResponseList(rules []recurringrule.RecurringRule) []RecurringRuleResponse {
	result := make([]RecurringRuleResponse, 0, len(rules))
	for i := range rules {
		result = append(result, toRecurringRuleResponse(&rules[i]))
	}
	return result
}
