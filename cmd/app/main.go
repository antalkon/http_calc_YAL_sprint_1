package main

import (
	"fmt"
	"net/http"

	"github.com/antalkon/http_calc_YAL_sprint_1/internal/router"
)

func main() {
	// Создаем маршрутизатор
	r := router.NewRouter()

	// Запуск сервера
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
