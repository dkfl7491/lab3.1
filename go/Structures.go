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
