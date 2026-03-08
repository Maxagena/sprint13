package api

import (
	"encoding/json"
	"net/http"

	"github.com/Maxagena/sprint13/pkg/db"
)

// addTaskHandler обрабатывает POST-запросы для добавления задач
func AddTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Invalid JSON"})
		return
	}

	if task.Title == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "Title is required"})
		return
	}

	if err := checkDate(&task); err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	id, err := db.AddTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "Failed to add task"})
		return
	}

	writeJson(w, http.StatusOK, map[string]string{"id": intToString(id)})
}
