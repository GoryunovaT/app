package main

import (
	"fmt"

	"github.com/GoryunovaT/app/calendar"
	"github.com/GoryunovaT/app/events"
)

func main() {
	e, err := events.NewEvent("Встреча", "2025/06/12 16:33")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	calendar.AddEvent("event1", e)

	e, err = events.NewEvent("Запись к окулисту", "2025/06/25 13:00")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	calendar.AddEvent("event2", e)

	e, err = events.NewEvent("Работа", "2025/06/16 08:00")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	calendar.AddEvent("event3", e)

	calendar.ShowEvents() // вывод всех событий
}
