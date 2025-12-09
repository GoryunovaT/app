package calendar

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/GoryunovaT/app/events"
	"github.com/GoryunovaT/app/storage"
)

type Calendar struct {
	calendarEvents map[string]*events.Event
	storage        storage.Storage
}

func NewCalendar(storage storage.Storage) *Calendar {
	if storage == nil {
		fmt.Println("ошибка: не удалось создать хранилище")
	}
	return &Calendar{
		calendarEvents: make(map[string]*events.Event),
		storage:        storage,
	}
}

func (c *Calendar) Save() error {
	if c.storage == nil {
		return errors.New("error")
	}
	data, err := json.Marshal(c.calendarEvents)
	if err != nil {
		return errors.New("ошибка сериализации")
	}
	return c.storage.Save(data)
}

func (c *Calendar) Load() error {
	data, err := c.storage.Load()
	if err != nil {
		return errors.New("ошибка десериализации")
	}
	return json.Unmarshal(data, &c.calendarEvents)
}

var EventsMap = make(map[string]*events.Event)

func (c *Calendar) AddEvent(title string, date string, priority events.Priority) (*events.Event, error) {
	e, err := events.NewEvent(title, date, priority)
	if err != nil {
		return nil, err
	}
	EventsMap[e.ID] = e
	fmt.Println("Событие добавлено:", e.Title)
	return e, nil
}

func (c *Calendar) ShowEvents() {
	for _, event := range EventsMap {
		fmt.Println(
			event.Title,
			"-",
			event.StartAt.Format("2006-01-02 15:04"))
	}
}

func (c *Calendar) DeleteEvent(id string) error {
	e, exist := EventsMap[id]
	if !exist {
		return errors.New("событие не найдено")
	}
	delete(EventsMap, id)
	fmt.Println("Событие", e.Title, "удалено")
	return nil
}

func (c *Calendar) EditEvent(id string, titleNew string, dateStrNew string, priority events.Priority) error {
	e, exists := EventsMap[id]
	if !exists {
		return errors.New("событие не найдено")
	}
	err := e.Update(titleNew, dateStrNew, priority)
	if err != nil {
		return err
	}
	fmt.Println("Событие изменено на", titleNew)
	return nil
}

func (c *Calendar) SetEventReminder(eventID string, message string, at time.Time) error {
	event, exist := EventsMap[eventID]
	if !exist {
		return errors.New("событие не найдено")
	}

	err := event.AddReminder(message, at)
	if err != nil {
		return errors.New("не удалось установить напоминание")
	}
	fmt.Println("напоминание установлено")
	return nil
}
