package reminder

import (
	"errors"
	"fmt"
	"time"
)

var allReminders []*Reminder

type Reminder struct {
	Message string
	At      time.Time
	Sent    bool
}

func NewReminder(message string, at time.Time) (*Reminder, error) {
	if message == "" {
		return nil, errors.New("сообщение не может быть пустым")
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

func DeleteReminder(message string) error {
	for i, reminder := range allReminders {
		if reminder.Message == message {
			if reminder.Sent {
				return errors.New("нельзя удалить уже отправленное напоминание")
			}
			allReminders = append(allReminders[:i], allReminders[i+1:]...)
			return nil
		}
	}
	return errors.New("напоминание не найдено")
}

func UpdateReminder(oldMessage, newMessage string, newTime time.Time) error {
	for _, reminder := range allReminders {
		if reminder.Message == oldMessage {
			if reminder.Sent {
				return errors.New("нельзя изменить уже отправленное напоминание")
			}
			if newMessage != "" {
				reminder.Message = newMessage
			}
			if !newTime.IsZero() {
				if newTime.Before(time.Now()) {
					return errors.New("время не может быть в прошлом")
				}
				reminder.At = newTime
			}
			return nil
		}
	}
	return errors.New("напоминание не найдено")
}

func ShowReminder(message string) error {
	for _, reminder := range allReminders {
		if reminder.Message == message {
			fmt.Println("Сообщение:", reminder.Message)
			fmt.Println("запланировано на:", reminder.At.Format("15:04 02.01.2006"))
			if reminder.Sent {
				fmt.Println("Уже отправлено")
			} else {
				fmt.Println("Еще не отправлено")
			}
			return nil
		}
	}
	return errors.New("напоминание не найдено")
}

func (r *Reminder) Send() error {
	if r.Sent {
		return errors.New("напоминание уже было отправлено")
	}
	fmt.Println("Напоминание!", r.Message)
	r.Sent = true
	return nil
}

func (r *Reminder) Stop() error {
	if r.Sent {
		return errors.New("напоминание уже отправлено, нельзя остановить")
	}
	r.At = time.Time{}
	fmt.Println("Напоминие остановлено")
	return nil
}
