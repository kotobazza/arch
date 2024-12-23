package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {


	// 1.	Инициализация подключений к базам данных
	// 2.	Поиск предмета по году обучения и семестру обучения
	// 3.	Поиск лекций, которые относятся к этому предмету
	// 4.	Поиск групп студентов, которые должны быть на этих лекциях согласно расписанию
	// 5.	Расчет общего числа студентов всех групп, которые должны быть на данной лекции
	// 6.	Формирование выходных данных и отображение пользователю



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
