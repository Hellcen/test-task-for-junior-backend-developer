package handlers

import (
	"errors"
	"example.com/taskservice/internal/domain/recurringrule"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func (h *RecurringHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	rule, err := h.usecase.GetByID(r.Context(), id)
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toRecurringRuleResponse(rule))
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

func (h *RecurringHandler) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	if err := h.usecase.Delete(r.Context(), id); err != nil {
		writeUsecaseError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *RecurringHandler) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var req RecurringRuleUpdateRequest
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	updated, err := h.usecase.Update(r.Context(), id, toUpdateInputDTO(req))
	if err != nil {
		writeUsecaseError(w, err)
		return
	}

	writeJSON(w, http.StatusOK, toRecurringRuleResponse(updated))
}

func (h *RecurringHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, errors.New("invalid id"))
		return
	}

	var req struct {
		Status recurringrule.Status `json:"status"`
	}
	if err := decodeJSON(r, &req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.usecase.UpdateStatus(r.Context(), id, req.Status); err != nil {
		writeUsecaseError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
