package transporthttp

import (
	"example.com/taskservice/internal/transport/http/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"

	"github.com/gorilla/mux"

	swaggerdocs "example.com/taskservice/internal/transport/http/docs"
	httphandlers "example.com/taskservice/internal/transport/http/handlers"
)

func NewRouter(
	taskHandler *httphandlers.TaskHandler,
	recurringRuleHandler *httphandlers.RecurringHandler,
	docsHandler *swaggerdocs.Handler,
) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)

	// Health check
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	}).Methods(http.MethodGet)

	router.HandleFunc("/swagger/openapi.json", docsHandler.ServeSpec).Methods(http.MethodGet)
	router.HandleFunc("/swagger/", docsHandler.ServeUI).Methods(http.MethodGet)
	router.HandleFunc("/swagger", docsHandler.RedirectToUI).Methods(http.MethodGet)

	api := router.PathPrefix("/api/v1").Subrouter()

	api.HandleFunc("/tasks", middleware.PromMiddleware(taskHandler.Create)).Methods(http.MethodPost)
	api.HandleFunc("/tasks", middleware.PromMiddleware(taskHandler.List)).Methods(http.MethodGet)
	api.HandleFunc("/tasks/{id:[0-9]+}", middleware.PromMiddleware(taskHandler.GetByID)).Methods(http.MethodGet)
	api.HandleFunc("/tasks/{id:[0-9]+}", middleware.PromMiddleware(taskHandler.Update)).Methods(http.MethodPut)
	api.HandleFunc("/tasks/{id:[0-9]+}", middleware.PromMiddleware(taskHandler.Delete)).Methods(http.MethodDelete)

	//Recurring handlers
	api.HandleFunc("/recurringrule-rules", middleware.PromMiddleware(recurringRuleHandler.Create)).Methods(http.MethodPost)
	api.HandleFunc("/recurringrule-rules", middleware.PromMiddleware(recurringRuleHandler.List)).Methods(http.MethodGet)
	api.HandleFunc("/recurringrule-rules/{id:[0-9]+}", middleware.PromMiddleware(recurringRuleHandler.GetByID)).Methods(http.MethodGet)
	api.HandleFunc("/recurringrule-rules/{id:[0-9]+}", middleware.PromMiddleware(recurringRuleHandler.Delete)).Methods(http.MethodDelete)
	api.HandleFunc("/recurringrule-rules/{id:[0-9]+}", middleware.PromMiddleware(recurringRuleHandler.Update)).Methods(http.MethodPut)
	api.HandleFunc("/recurringrule-rules/{id:[0-9]+}/status", middleware.PromMiddleware(recurringRuleHandler.UpdateStatus)).Methods(http.MethodPatch)

	return router
}
