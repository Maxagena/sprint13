package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Maxagena/sprint13/pkg/db"
)

// TasksResp структура для ответа со списком задач
type TasksResp struct {
	Tasks []*db.Task `json:"tasks"`
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]string{"error": "Method not allowed"})
		return
	}

	search := r.URL.Query().Get("search")
	tasks, err := db.Tasks(50, search)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": "internal server error"})
		return
	}

	if tasks == nil {
		tasks = []*db.Task{}
	}

	writeJson(w, http.StatusOK, TasksResp{Tasks: tasks})
}

// GetTaskHandler обработчик GET /api/task
func GetTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "не указан идентификатор"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	writeJson(w, http.StatusOK, task)
}

// UpdateTaskHandler обработчик PUT /api/task
func UpdateTaskHandler(w http.ResponseWriter, r *http.Request) {
	var task db.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "ошибка десериализации JSON"})
		return
	}

	if task.ID == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "не указан идентификатор задачи"})
		return
	}

	if task.Title == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "не указан заголовок задачи"})
		return
	}

	err = checkDate(&task)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = db.UpdateTask(&task)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	// Успешный ответ — пустой JSON
	writeJson(w, http.StatusOK, map[string]interface{}{})
}

func DeleteTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{
			"error": "не указан идентификатор",
		})
		return
	}

	if err := db.DeleteTask(id); err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]interface{}{})
}

func DoneTaskHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": "не указан идентификатор"})
		return
	}

	task, err := db.GetTask(id)
	if err != nil {
		writeJson(w, http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}

	if task.Repeat == "" {
		// Одноразовая задача — удаляем
		err := db.DeleteTask(id)
		if err != nil {
			writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		writeJson(w, http.StatusOK, map[string]interface{}{})
		return
	}

	// Периодическая задача — вычисляем следующую дату
	nextDate, err := NextDate(time.Now(), task.Date, task.Repeat)
	if err != nil {
		writeJson(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	err = db.UpdateDate(nextDate, id)
	if err != nil {
		writeJson(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJson(w, http.StatusOK, map[string]interface{}{})
}
