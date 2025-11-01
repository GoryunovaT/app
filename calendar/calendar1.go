package calendar

import (
	"events" // импортируем пакет с Событиями
	"fmt"
)

var eventsMap = make(map[string]events.Event) // мапа событий

func AddEvent(key string, e events.Event) { // принимаем события в аргументе
	eventsMap[key] = e                         // добавляем события по ключу
	fmt.Println("Событие добавлено:", e.Title) // выводим лог для проверки
}
