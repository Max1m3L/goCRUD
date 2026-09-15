package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	connections := []string{
		"user=postgres password=root dbname=cosmic_db sslmode=disable host=localhost port=5432",
		"user=postgres password=root dbname=cosmic_db sslmode=disable host=127.0.0.1 port=5432",
		"user=postgres password=root dbname=cosmic_db sslmode=disable host=::1 port=5432",
		"user=postgres password=root dbname=cosmic_db sslmode=disable host=localhost port=5433",
		"user=postgres password=root dbname=cosmic_db sslmode=disable host=127.0.0.1 port=5433",
	}

	for i, connStr := range connections {
		fmt.Printf("Попытка %d: %s\n", i+1, connStr)

		db, err := sql.Open("postgres", connStr)
		if err != nil {
			fmt.Printf("  ❌ Ошибка: %v\n\n", err)
			continue
		}

		err = db.Ping()
		if err != nil {
			fmt.Printf("  ❌ Ошибка пинга: %v\n\n", err)
			db.Close()
			continue
		}

		fmt.Println("  ✅ УСПЕШНОЕ ПОДКЛЮЧЕНИЕ!")
		fmt.Printf("  Работает на: %s\n\n", connStr)
		db.Close()
		return
	}

	fmt.Println("❌ Все варианты не сработали")
	fmt.Println("\n💡 Проверь:")
	fmt.Println("1. Запущен ли PostgreSQL (Службы -> postgresql)")
	fmt.Println("2. Какой порт использует (в pgAdmin -> Properties)")
	fmt.Println("3. Нет ли конфликта с Docker")
}
