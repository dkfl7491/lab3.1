package main

import "errors"

// Массив

type ArrayADT struct {
    Data []string
}

// Добавление элемента в конец массива.
func (a *ArrayADT) Push(value string) {
    a.Data = append(a.Data, value)
}

// Вставка элемента по индексу.
func (a *ArrayADT) Insert(index int, value string) error {
    if index < 0 || index > len(a.Data) {
        return errors.New("invalid array index")
    }

    a.Data = append(a.Data, "")
    copy(a.Data[index+1:], a.Data[index:])
    a.Data[index] = value
    return nil
}

// Получение элемента по индексу.
func (a *ArrayADT) Get(index int) (string, error) {
    if index < 0 || index >= len(a.Data) {
        return "", errors.New("invalid array index")
    }

    return a.Data[index], nil
}

// Удаление элемента по индексу.
func (a *ArrayADT) Remove(index int) error {
    if index < 0 || index >= len(a.Data) {
        return errors.New("invalid array index")
    }

    a.Data = append(a.Data[:index], a.Data[index+1:]...)
    return nil
}

// Замена значения по индексу.
func (a *ArrayADT) Set(index int, value string) error {
    if index < 0 || index >= len(a.Data) {
        return errors.New("invalid array index")
    }

    a.Data[index] = value
    return nil
}

// Копия элементов массива.
func (a *ArrayADT) Values() []string {
    return append([]string(nil), a.Data...)
}
