package reminder

import (
	"errors"
	"fmt"
	"time"
)

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
}

func NewReminder(message string, at time.Time) (*Reminder, error) {
	if !IsValidTitle(message) {
		return nil, errors.New("некорректное сообщение")
	}
	if at.IsZero() {
		return nil, errors.New("время не указано")
	}
	if at.Before(time.Now()) {
		return nil, errors.New("время напоминания не может быть в прошлом")
	}

	return &Reminder{
		Message: message,
		At:      at,
		Sent:    false,
	}, nil
	// Позже здесь будет код который следит за временем и запускает напоминание
	// типа registerReminder(r) или time.AfterFunc()
}

func DeleteReminder() error {

}

func UpdateReminder() error {

}

func ShowReminder() error {

}
func (r *Reminder) Send() error {
	if r.Sent {
		return errors.New("напоминание уже было отправлено")
	}
	fmt.Println("Reminder!", r.Message)
	r.Sent = true
	return nil
}

func (r *Reminder) Stop() {
	if r.Sent {
		fmt.Println("Напоминание уже отправлено, нельзя остановить")
	}
	r.At = time.Time{}
	fmt.Println("Напоминие остановлено")
}
