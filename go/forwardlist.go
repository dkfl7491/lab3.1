package main

import "errors"

// Односвязный список

type SNode struct {
    Value string
    Next  *SNode
}

type SinglyList struct {
    Head  *SNode
    Tail  *SNode
    Count int
}

// Добавление элемента в конец списка.
func (l *SinglyList) PushBack(value string) {
    node := &SNode{Value: value}

    if l.Head == nil {
        l.Head = node
    } else {
        l.Tail.Next = node
    }

    l.Tail = node
    l.Count++
}

// Добавление элемента в начало списка.
func (l *SinglyList) PushFront(value string) {
    node := &SNode{Value: value, Next: l.Head}
    l.Head = node

    if l.Tail == nil {
        l.Tail = node
    }

    l.Count++
}

// Вставка элемента перед указанным значением.
func (l *SinglyList) InsertBefore(target, value string) error {
    if l.Head == nil {
        return errors.New("list is empty")
    }

    if l.Head.Value == target {
        l.PushFront(value)
        return nil
    }

    previous := l.Head

    for previous.Next != nil && previous.Next.Value != target {
        previous = previous.Next
    }

    if previous.Next == nil {
        return errors.New("target value not found")
    }

    previous.Next = &SNode{
        Value: value,
        Next:  previous.Next,
    }

    l.Count++
    return nil
}

// Вставка элемента после указанного значения.
func (l *SinglyList) InsertAfter(target, value string) error {
    current := l.Head

    for current != nil && current.Value != target {
        current = current.Next
    }

    if current == nil {
        return errors.New("target value not found")
    }

    node := &SNode{
        Value: value,
        Next:  current.Next,
    }

    current.Next = node

    if l.Tail == current {
        l.Tail = node
    }

    l.Count++
    return nil
}

// Удаление первого элемента списка.
func (l *SinglyList) RemoveFront() error {
    if l.Head == nil {
        return errors.New("list is empty")
    }

    l.Head = l.Head.Next
    l.Count--

    if l.Head == nil {
        l.Tail = nil
    }

    return nil
}

// Удаление последнего элемента списка.
func (l *SinglyList) RemoveBack() error {
    if l.Head == nil {
        return errors.New("list is empty")
    }

    if l.Head == l.Tail {
        l.Head = nil
        l.Tail = nil
        l.Count = 0
        return nil
    }

    current := l.Head

    for current.Next != l.Tail {
        current = current.Next
    }

    current.Next = nil
    l.Tail = current
    l.Count--
    return nil
}

// Удаление первого найденного значения.
func (l *SinglyList) RemoveValue(value string) error {
    if l.Head == nil {
        return errors.New("list is empty")
    }

    if l.Head.Value == value {
        return l.RemoveFront()
    }

    previous := l.Head

    for previous.Next != nil && previous.Next.Value != value {
        previous = previous.Next
    }

    if previous.Next == nil {
        return errors.New("value not found")
    }

    if previous.Next == l.Tail {
        l.Tail = previous
    }

    previous.Next = previous.Next.Next
    l.Count--
    return nil
}

// Удаление элемента непосредственно перед указанным значением.
func (l *SinglyList) RemoveBefore(target string) error {
    if l.Head == nil || l.Head.Next == nil {
        return errors.New("element before target not found")
    }

    if l.Head.Next.Value == target {
        l.Head = l.Head.Next
        l.Count--
        return nil
    }

    previous := l.Head
    current := l.Head.Next

    for current.Next != nil && current.Next.Value != target {
        previous = current
        current = current.Next
    }

    if current.Next == nil {
        return errors.New("target value not found or has no previous element")
    }

    previous.Next = current.Next
    l.Count--
    return nil
}

// Удаление элемента непосредственно после указанного значения.
func (l *SinglyList) RemoveAfter(target string) error {
    current := l.Head

    for current != nil && current.Value != target {
        current = current.Next
    }

    if current == nil {
        return errors.New("target value not found")
    }

    if current.Next == nil {
        return errors.New("element after target not found")
    }

    if current.Next == l.Tail {
        l.Tail = current
    }

    current.Next = current.Next.Next
    l.Count--
    return nil
}

// Проверка наличия значения.
func (l *SinglyList) Contains(value string) bool {
    for current := l.Head; current != nil; current = current.Next {
        if current.Value == value {
            return true
        }
    }

    return false
}

// Получение элемента по индексу.
func (l *SinglyList) Get(index int) (string, error) {
    if index < 0 || index >= l.Count {
        return "", errors.New("invalid list index")
    }

    current := l.Head

    for i := 0; i < index; i++ {
        current = current.Next
    }

    return current.Value, nil
}

// Последовательность значений списка.
func (l *SinglyList) Values() []string {
    result := make([]string, 0, l.Count)

    for current := l.Head; current != nil; current = current.Next {
        result = append(result, current.Value)
    }

    return result
}
