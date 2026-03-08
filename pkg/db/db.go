package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "modernc.org/sqlite"
)

var db *sql.DB

const schema = `
CREATE TABLE scheduler (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    date CHAR(8) NOT NULL DEFAULT "",
    title VARCHAR(255) NOT NULL DEFAULT "",
    comment TEXT NOT NULL DEFAULT "",
    repeat VARCHAR(128) NOT NULL DEFAULT ""
);
CREATE INDEX idx_scheduler_date ON scheduler(date);
`

// Init открывает базу данных и создает таблицу при необходимости
func Init(dbFile string) error {
	install := false

	// Проверяем наличие файла базы данных
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		install = true
	}

	// Открываем подключение к базе
	var err error
	db, err = sql.Open("sqlite", dbFile)
	if err != nil {
		return fmt.Errorf("database opening error: %w", err)
	}

	// Если файл не существовал создаем таблицу и индекс
	if install {
		if _, err = db.Exec(schema); err != nil {
			return fmt.Errorf("schema creation error: %w", err)
		}

	}

	return nil
}

// Close закрывает соединение с базой данных
func GetDB() *sql.DB {
	return db
}
