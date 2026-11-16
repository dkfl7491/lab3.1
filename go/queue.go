package main

import "errors"

// Очередь

type QueueADT struct {
    Data []string
}

// Добавление элемента в конец очереди.
func (q *QueueADT) Push(value string) {
    q.Data = append(q.Data, value)
}

// Извлечение первого элемента очереди.
func (q *QueueADT) Pop() (string, error) {
    if len(q.Data) == 0 {
        return "", errors.New("queue is empty")
    }

    value := q.Data[0]
    q.Data = q.Data[1:]
    return value, nil
}

// Просмотр первого элемента без удаления.
func (q *QueueADT) Get() (string, error) {
    if len(q.Data) == 0 {
        return "", errors.New("queue is empty")
    }

    return q.Data[0], nil
}

// Копия элементов очереди.
func (q *QueueADT) Values() []string {
    return append([]string(nil), q.Data...)
}
