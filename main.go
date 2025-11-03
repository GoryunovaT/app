package main

import (
	"fmt"

	"github.com/GoryunovaT/app/calendar"
)

func main() {
	_, err := calendar.AddEvent("Встреча", "2025/06/12 16:33")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	e2, err := calendar.AddEvent("Заказать роллы", "2025/03/12 20:43")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	_, err = calendar.AddEvent("Помыть пол", "2025/06/25 17:50")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	e4, err := calendar.AddEvent("Покормить собаку", "2025/06/02 02:33")
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	calendar.ShowEvents()

	calendar.DeleteEvent(e2.ID)
	calendar.ShowEvents()

	calendar.EditEvent(e4.ID, "Перевести будильник", "2025/07/14 07:20")
	calendar.ShowEvents()

}
