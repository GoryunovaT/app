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
	if t.Before(time.Now()) {
		return nil, errors.New("дата события не может быть в прошлом")
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
	if newDate.Before(time.Now()) {
		return errors.New("новая дата не может быть в прошлом")
	}

	e.Title = title
	e.StartAt = newDate
	return nil
}

func (e *Event) AddReminder(message string, at time.Time) error {
	if at.After(e.StartAt) {
		return errors.New("напоминание не может быть позже события")
	}
	if at.Before(time.Now()) {
		return errors.New("время напоминания не может быть в прошлом")
	}
	r, err := reminder.NewReminder(message, at)
	if err != nil {
		return errors.New("ошибка создания напоминания")
	}
	e.Reminder = r
	return nil
}
func (e *Event) RemoveReminder() error {
	if e.Reminder == nil {
		return errors.New("у события нет напоминания")
	}
	e.Reminder = nil
	return nil
}
