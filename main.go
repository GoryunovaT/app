package main

import (
	"app/calendar"
	"app/events"
	"fmt"
	"time"
)

func main() {
	e := events.Event{ // создаем событие типа Event
		Title:   "Встреча",
		StartAt: time.Now(), // простая работа с датой :)
	}
	calendar.AddEvent("event1", e)    // добавляем событие в календарь
	fmt.Println("Календарь обновлён") // сообщаем о результате
}
