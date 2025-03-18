package database

import (
	"github.com/jmoiron/sqlx"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sqlx.DB

func InitDB() {
	var err error
	DB, err = sqlx.Connect("sqlite", "schedule.db")
	if err != nil {
		log.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Создаём таблицу, если её нет
	schema := `
	CREATE TABLE IF NOT EXISTS schedule (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		day TEXT UNIQUE,
		subject TEXT
	);
	`
	DB.MustExec(schema)
	log.Println("База данных инициализирована")
}

func GetSchedule() map[string]string {
	schedule := make(map[string]string)
	rows, err := DB.Query("SELECT day, subject FROM schedule")
	if err != nil {
		log.Println("Ошибка при запросе расписания:", err)
		return schedule
	}
	defer rows.Close()

	for rows.Next() {
		var day, subject string
		if err := rows.Scan(&day, &subject); err == nil {
			schedule[day] = subject
		}
	}
	return schedule
}

func AddSchedule(day, subject string) {
	_, err := DB.Exec("INSERT INTO schedule (day, subject) VALUES (?, ?) ON CONFLICT(day) DO UPDATE SET subject = excluded.subject", day, subject)
	if err != nil {
		log.Println("Ошибка при добавлении занятия:", err)
	}
}
