package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// Твой родной JDBC-формат, но теперь через Docker
	connStr := "user=postgres password=root dbname=cosmic_db sslmode=disable host=localhost port=5432"

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Ошибка подключения к БД:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Не удалось пинговать БД:", err)
	}

	fmt.Println("✅ Подключено к PostgreSQL через Docker!")
	createTable()
}

func createTable() {
	query := `
    CREATE TABLE IF NOT EXISTS spaceships (
        id SERIAL PRIMARY KEY,
        name VARCHAR(100) NOT NULL UNIQUE,
        fuel INT DEFAULT 0,
        is_ready BOOLEAN DEFAULT FALSE,
        engine_type VARCHAR(50),
        crew_count INT DEFAULT 0
    );`

	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Ошибка создания таблицы:", err)
	}
	fmt.Println("✅ Таблица spaceships создана/существует")
}
