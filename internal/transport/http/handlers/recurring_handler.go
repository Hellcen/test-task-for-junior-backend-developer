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
