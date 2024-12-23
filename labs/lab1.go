package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	// Настройки подключения к базе данных
	connStr := "user=username dbname=mydb password=mypassword host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Проверка подключения
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// Выполнение запроса
	rows, err := db.Query("SELECT id, name FROM mytable")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Обработка результатов
	for rows.Next() {
		var id int
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s\n", id, name)
	}

	// Проверка на ошибки после завершения обработки
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
