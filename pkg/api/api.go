package api

import (
	"net/http"
)

// Init регистрирует все API обработчики
func Init() {
	http.HandleFunc("/api/nextdate", (NextDateHandler))
	http.HandleFunc("/api/task", taskHandler)
	http.HandleFunc("/api/tasks", tasksHandler)
	http.HandleFunc("/api/task/done", DoneTaskHandler)

}
