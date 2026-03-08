package api

import (
	"net/http"
)

// taskHandler обрабатывает различные HTTP методы для /api/task
func taskHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		AddTaskHandler(w, r)
	case http.MethodGet:
		GetTaskHandler(w, r)
	case http.MethodPut:
		UpdateTaskHandler(w, r)
	case http.MethodDelete:
		DeleteTaskHandler(w, r)
	default:
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{
			"error": "Method not allowed",
		})
	}
}
