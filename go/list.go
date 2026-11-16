package main

import "errors"

// Двусвязный список

type DNode struct {
    Value string
    Prev  *DNode
    Next  *DNode
}

type DoublyList struct {
    Head  *DNode
    Tail  *DNode
    Count int
}

// Добавление элемента в конец списка.
func (l *DoublyList) PushBack(value string) {
    node := &DNode{Value: value, Prev: l.Tail}

    if l.Tail != nil {
        l.Tail.Next = node
    } else {
        l.Head = node
    }

    l.Tail = node
    l.Count++
}

// Добавление элемента в начало списка.
func (l *DoublyList) PushFront(value string) {
    node := &DNode{Value: value, Next: l.Head}

    if l.Head != nil {
        l.Head.Prev = node
    } else {
        l.Tail = node
    }

    l.Head = node
    l.Count++
}

// Поиск узла по значению.
func (l *DoublyList) find(value string) *DNode {
    for current := l.Head; current != nil; current = current.Next {
        if current.Value == value {
            return current
        }
    }

    return nil
}

// Вставка элемента перед указанным значением.
func (l *DoublyList) InsertBefore(target, value string) error {
    current := l.find(target)

    if current == nil {
        return errors.New("target value not found")
    }

    if current == l.Head {
        l.PushFront(value)
        return nil
    }

    node := &DNode{
        Value: value,
        Prev:  current.Prev,
        Next:  current,
    }

    current.Prev.Next = node
    current.Prev = node
    l.Count++
    return nil
}

// Вставка элемента после указанного значения.
func (l *DoublyList) InsertAfter(target, value string) error {
    current := l.find(target)

    if current == nil {
        return errors.New("target value not found")
    }

    if current == l.Tail {
        l.PushBack(value)
        return nil
    }

    node := &DNode{
        Value: value,
        Prev:  current,
        Next:  current.Next,
    }

    current.Next.Prev = node
    current.Next = node
    l.Count++
    return nil
}

// Удаление узла с обновлением соседних связей.
func (l *DoublyList) removeNode(node *DNode) {
    if node.Prev != nil {
        node.Prev.Next = node.Next
    } else {
        l.Head = node.Next
    }

    if node.Next != nil {
        node.Next.Prev = node.Prev
    } else {
        l.Tail = node.Prev
    }

    l.Count--
}

// Удаление элемента непосредственно перед указанным значением.
func (l *DoublyList) RemoveBefore(target string) error {
    current := l.find(target)

    if current == nil {
        return errors.New("target value not found")
    }

    if current.Prev == nil {
        return errors.New("element before target not found")
    }

    l.removeNode(current.Prev)
    return nil
}

// Удаление элемента непосредственно после указанного значения.
func (l *DoublyList) RemoveAfter(target string) error {
    current := l.find(target)

    if current == nil {
        return errors.New("target value not found")
    }

    if current.Next == nil {
        return errors.New("element after target not found")
    }

    l.removeNode(current.Next)
    return nil
}

// Удаление первого элемента списка.
func (l *DoublyList) RemoveFront() error {
    if l.Head == nil {
        return errors.New("list is empty")
    }

    l.removeNode(l.Head)
    return nil
}

// Удаление последнего элемента списка.
func (l *DoublyList) RemoveBack() error {
    if l.Tail == nil {
        return errors.New("list is empty")
    }

    l.removeNode(l.Tail)
    return nil
}

// Удаление элемента по значению.
func (l *DoublyList) RemoveValue(value string) error {
    node := l.find(value)

    if node == nil {
        return errors.New("value not found")
    }

    l.removeNode(node)
    return nil
}

// Проверка наличия значения.
func (l *DoublyList) Contains(value string) bool {
    return l.find(value) != nil
}

// Получение элемента по индексу.
func (l *DoublyList) Get(index int) (string, error) {
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
func (l *DoublyList) Values() []string {
    result := make([]string, 0, l.Count)

    for current := l.Head; current != nil; current = current.Next {
        result = append(result, current.Value)
    }

    return result
}
