package main

import (
	"fmt"

	"github.com/GoryunovaT/app/calendar"
	"github.com/GoryunovaT/app/events"
	"github.com/GoryunovaT/app/storage"
)

func main() {
	s := storage.ZipStorage("my_calendar.zip")

	myCalendar := calendar.NewCalendar(s)

	defer func() {
		err := myCalendar.Save()
		if err != nil {
			fmt.Println("ошибка при сохранении:", err)
		}
	}()

	_, err := myCalendar.AddEvent("Встреча", "2025/06/12 16:33", events.PriorityHigh)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	e2, err := myCalendar.AddEvent("Заказать роллы", "2025/03/12 20:43", events.PriorityMedium)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	_, err = myCalendar.AddEvent("Помыть пол", "2025/06/25 17:50", events.PriorityLow)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	e4, err := myCalendar.AddEvent("Покормить кота", "2025/06/02 02:33", events.PriorityMedium)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}
	myCalendar.ShowEvents()

	myCalendar.DeleteEvent(e2.ID)
	myCalendar.ShowEvents()

	myCalendar.EditEvent(e4.ID, "Перевести будильник", "2025/07/14 07:20", events.PriorityHigh)
	myCalendar.ShowEvents()

}
