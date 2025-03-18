package main

import "my_bot/database"

func main() {
	database.InitDB()
	database.AddSchedule("Понедельник", "Основы робототехники")
	database.AddSchedule("Вторник", "Программирование")
	database.AddSchedule("Среда", "Механика")
	database.AddSchedule("Четверг", "Практика")
	database.AddSchedule("Пятница", "Подготовка к соревнованиям")
}
