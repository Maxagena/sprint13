package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Maxagena/sprint13/pkg/db"
)

const dateFormat = "20060102"

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func intToString(id int64) string {
	return fmt.Sprintf("%d", id)
}

func checkDate(task *db.Task) error {
	now := time.Now()
	// если пустая дата — ставим сегодня
	if task.Date == "" {
		task.Date = now.Format(dateFormat)
		return nil
	}

	// парсим указанную дату
	t, err := time.ParseInLocation(dateFormat, task.Date, now.Location())
	if err != nil {
		return fmt.Errorf("invalid date format: %w", err)
	}

	// если есть правило повторения — проверяем его через NextDate
	var next string
	if task.Repeat != "" {
		next, err = NextDate(now, task.Date, task.Repeat)
		if err != nil {
			return fmt.Errorf("invalid repeat: %w", err)
		}
	}

	// Если указанная дата раньше или равна сегодня (now), корректируем:
	if afterNow(now, t) {
		if task.Repeat == "" {
			// без повтора — ставим сегодняшнюю дату
			task.Date = now.Format(dateFormat)
		} else {
			// с повтором — ставим вычисленную следующую дату
			task.Date = next
		}
	}
	return nil
}
