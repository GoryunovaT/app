package calendar

import (
	"errors"
	"fmt"

	"github.com/GoryunovaT/app/events"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(key string, e events.Event) error {
	if !IsValidTitle(e.Title) {
		return errors.New("ошибка валидации")
	}
	eventsMap[key] = e
	fmt.Println("Событие добавлено:", e.Title)
	return nil
}

func ShowEvents() {
	for _, event := range eventsMap {
		fmt.Println(
			event.Title,
			"-",
			event.StartAt.Format("2006-01-02 15:04"))
	}
}
