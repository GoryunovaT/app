package events

import (
	"errors"
	"fmt"
	"time"

	"github.com/GoryunovaT/app/reminder"
	"github.com/araddon/dateparse"
	"github.com/google/uuid"
)

type Event struct {
	ID       string
	Title    string
	StartAt  time.Time
	Priority Priority
	Reminder *reminder.Reminder
}

func (e *Event) Print() {
	fmt.Println(e.Title, e.StartAt)
}

func getNextID() string {
	return uuid.New().String()
}

func NewEvent(title string, dateStr string, priority Priority) (*Event, error) {
	if !IsValidTitle(title) {
		return nil, errors.New("некорректное наименование задачи")
	}
	if err := priority.Validate(); err != nil {
		return nil, errors.New("некорректный приоритет")
	}
	t, err := dateparse.ParseAny(dateStr)
	if err != nil {
		return nil, errors.New("неверный формат даты")
	}
	return &Event{
		ID:       getNextID(),
		Title:    title,
		StartAt:  t,
		Priority: priority,
		Reminder: nil,
	}, nil
}

func (e *Event) Update(title string, date string, priority Priority) error {
	if !IsValidTitle(title) {
		return errors.New("недопустимое название")
	}
	if err := priority.Validate(); err != nil {
		return errors.New("некорректный приоритет")
	}
	newDate, err := time.Parse("2006/01/02 15:04", date)
	if err != nil {
		return errors.New("неверный формат даты")
	}
	e.Title = title
	e.StartAt = newDate
	return nil
}
