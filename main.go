package main

import (
	"fmt"

	"github.com/GoryunovaT/app/calendar"
	"github.com/GoryunovaT/app/events"
	"github.com/GoryunovaT/app/storage"
)

func main() {
	s := storage.NewZipStorage("my_calendar.zip")

	myCalendar := calendar.NewCalendar(s)

	err := myCalendar.Load()
	if err != nil {
		fmt.Println("Не удалось загрузить сохраненные события:", err)
		fmt.Println("Создаем новый календарь")
	}

	defer func() {
		err := myCalendar.Save()
		if err != nil {
			fmt.Println("ошибка при сохранении:", err)
		} else {
			fmt.Println("Календарь сохранен")
		}
	}()

	_, err = myCalendar.AddEvent("Встреча", "2025/12/29 16:33", events.PriorityHigh)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	event2, err := myCalendar.AddEvent("Заказать роллы", "2025/12/30 20:43", events.PriorityLow)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	_, err = myCalendar.AddEvent("Помыть пол", "2025/12/31 17:50", events.PriorityMedium)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	event4, err := myCalendar.AddEvent("Покормить кота", "2025/12/18 02:33", events.PriorityMedium)
	if err != nil {
		fmt.Println("Ошибка создания события:", err)
		return
	}

	myCalendar.ShowEvents()

	myCalendar.DeleteEvent(event2.ID)
	myCalendar.ShowEvents()

	myCalendar.EditEvent(event4.ID, "Перевести будильник", "2026/01/01 07:20", events.PriorityHigh)
	myCalendar.ShowEvents()

}
