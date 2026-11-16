package main

import "errors"

// Стек

type StackADT struct {
    Data []string
}

// Добавление элемента на вершину стека.
func (s *StackADT) Push(value string) {
    s.Data = append(s.Data, value)
}

// Извлечение элемента с вершины стека.
func (s *StackADT) Pop() (string, error) {
    if len(s.Data) == 0 {
        return "", errors.New("stack is empty")
    }

    index := len(s.Data) - 1
    value := s.Data[index]
    s.Data = s.Data[:index]
    return value, nil
}

// Просмотр верхнего элемента без удаления.
func (s *StackADT) Get() (string, error) {
    if len(s.Data) == 0 {
        return "", errors.New("stack is empty")
    }

    return s.Data[len(s.Data)-1], nil
}

// Копия элементов стека.
func (s *StackADT) Values() []string {
    return append([]string(nil), s.Data...)
}
