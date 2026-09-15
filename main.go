package main

import (
	"cosmic-crud/config"
	"cosmic-crud/handlers"
	"cosmic-crud/repositories"
	"cosmic-crud/services"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.InitDB()
	db := config.DB
	defer db.Close()

	repo := repositories.NewSpaceshipRepo(db)
	service := services.NewSpaceshipService(repo)
	handler := handlers.NewSpaceshipHandler(service)

	// Регистрация маршрутов
	http.HandleFunc("/spaceships", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handler.GetAll(w, r)
		case http.MethodPost:
			handler.Create(w, r)
		default:
			http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/spaceships/", func(w http.ResponseWriter, r *http.Request) {
		// Проверяем, не /spaceships/refuel ли это
		if r.URL.Path == "/spaceships/refuel" {
			http.Error(w, "Неверный запрос", http.StatusBadRequest)
			return
		}

		switch r.Method {
		case http.MethodGet:
			handler.GetByID(w, r)
		case http.MethodPut:
			handler.Update(w, r)
		case http.MethodDelete:
			handler.Delete(w, r)
		default:
			http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/spaceships/refuel", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Метод не разрешён", http.StatusMethodNotAllowed)
			return
		}
		handler.Refuel(w, r)
	})

	// Запуск сервера
	fmt.Println("🚀 Сервер запущен на http://localhost:8080")
	fmt.Println("📋 Доступные эндпоинты:")
	fmt.Println("  GET    /spaceships           - список всех кораблей")
	fmt.Println("  POST   /spaceships           - создать корабль")
	fmt.Println("  GET    /spaceships/{id}      - получить корабль по ID")
	fmt.Println("  PUT    /spaceships/{id}      - обновить корабль")
	fmt.Println("  DELETE /spaceships/{id}      - удалить корабль")
	fmt.Println("  POST   /spaceships/{id}/refuel - заправить корабль")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
