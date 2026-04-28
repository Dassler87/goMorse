package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	// Создаём логгер
	logger := log.New(os.Stdout, "SERVER: ", log.LstdFlags)

	// Создаём сервер
	srv := server.New(logger)

	// Запускаем сервер
	logger.Println("Запуск сервера на :8080")
	err := srv.Server.ListenAndServe()
	if err != nil {
		logger.Fatal("Ошибка запуска сервера: ", err)
	}
}
