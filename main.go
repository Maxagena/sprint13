package main

import (
	"log"
	"os"

	"github.com/Maxagena/sprint13/pkg/db"
	"github.com/Maxagena/sprint13/pkg/server"
)

func main() {
	err := db.Init("scheduler.db")
	if err != nil {
		log.Fatalf("Ошибка инициализации базы данных: %v", err)
	}
	defer db.GetDB().Close()

	logger := log.New(os.Stdout, "server: ", log.LstdFlags)

	srv := server.NewServer(logger)

	logger.Println("Сервер запускается на http://localhost:7540 ...")
	if err := srv.HTTPServer.ListenAndServe(); err != nil {
		logger.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
