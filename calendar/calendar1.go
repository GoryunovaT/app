package calendar

import (
	"errors"
	"fmt"
	"time"

	"github.com/GoryunovaT/app/events"
)

var EventsMap = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {
	e, err := events.NewEvent(title, date)
	if err != nil {
		return events.Event{}, err
	}
	EventsMap[e.ID] = e
	fmt.Println("Событие добавлено:", e.Title)
	return e, nil
}

func ShowEvents() {
	for _, event := range EventsMap {
		fmt.Println(
			event.Title,
			"-",
			event.StartAt.Format("2006-01-02 15:04"))
	}
}

func FindEvents(id string) (events.Event, error) {
	event, exists := EventsMap[id]
	if !exists {
		return events.Event{}, errors.New("событие с таким ID не найдено")
	}
	return event, nil
}

func DeleteEvent(id string) error {
	foundEvent, err := FindEvents(id)
	if err != nil {
		return err
	}
	delete(EventsMap, id)
	fmt.Println("Событие", foundEvent.Title, "удалено")
	return nil
}

func EditEvent(id string, titleNew string, dateStrNew string) error {
	foundEvent, err := FindEvents(id)
	if err != nil {
		return err
	}
	if !events.IsValidTitle(titleNew) {
		return errors.New("недопустимое название")
	}
	newDate, err := time.Parse("2006/01/02 15:04", dateStrNew)
	if err != nil {
		return errors.New("неверный формат даты")
	}
	EventsMap[id] = events.Event{
		ID:      id,
		Title:   titleNew,
		StartAt: newDate,
	}
	fmt.Println("Событие", foundEvent.Title, "изменено на", titleNew)
	return nil
}
