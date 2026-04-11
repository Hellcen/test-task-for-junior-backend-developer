package handlers

import (
	"net/http"

	recurringusecase "example.com/taskservice/internal/usecase/recurring"
)

type RecurringHandler struct {
	usecase recurringusecase.Usecase
}

func NewRecurringHandler(usecase recurringusecase.Usecase) *RecurringHandler {
	return &RecurringHandler{usecase: usecase}
}

func (h *RecurringHandler) Create(w http.ResponseWriter, r *http.Request) {
	var req RecurringRuleCreateRequest

	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	created, err := h.usecase.Create(r.Context(), toCreateInputDTO(req))
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	writeJSON(w, http.StatusCreated, toRecurringRuleResponse(created))
}

func (h *RecurringHandler) List(w http.ResponseWriter, r *http.Request) {
	rules, err := h.usecase.List(r.Context())
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	response := make([]RecurringRuleResponse, 0, len(rules))
	for i := range rules {
		response = append(response, toRecurringRuleResponse(&rules[i]))
	}

	writeJSON(w, http.StatusOK, response)
}
