package main

import (
    "errors"
    "fmt"
    "strconv"
    "strings"
)

// Тип обработчика команды.
type Handler func([]string) error

// Данные обработчика команд и набор доступных операций.
type CommandProcessor struct {
    DB       *Database
    Filename string
    Handlers map[string]Handler
}

// Инициализация обработчика и регистрация команд.
func NewCommandProcessor(db *Database, filename string) *CommandProcessor {
    processor := &CommandProcessor{
        DB:       db,
        Filename: filename,
        Handlers: make(map[string]Handler),
    }

    processor.registerHandlers()
    return processor
}

// Проверка количества аргументов команды.
func checkArgs(args []string, count int) error {
    if len(args) != count {
        return errors.New("incorrect number of arguments")
    }

    return nil
}

// Преобразование индекса и проверка его корректности.
func indexOf(value string) (int, error) {
    index, err := strconv.Atoi(value)

    if err != nil || index < 0 {
        return 0, errors.New("invalid index")
    }

    return index, nil
}

// Вывод последовательности значений.
func printValues(values []string) {
    fmt.Println(strings.Join(values, " "))
}

// Регистрация обработчиков операций для всех структур.
func (p *CommandProcessor) registerHandlers() {
    h := p.Handlers
    db := p.DB

    // Операции с массивом

    h["MPUSH"] = func(a []string) error {
        if err := checkArgs(a, 2); err != nil {
            return err
        }

        if db.Arrays[a[0]] == nil {
            db.Arrays[a[0]] = &ArrayADT{}
        }

        db.Arrays[a[0]].Push(a[1])
        return nil
    }

    h["MINSERT"] = func(a []string) error {
        if err := checkArgs(a, 3); err != nil {
            return err
        }

        index, err := indexOf(a[1])
        if err != nil {
            return err
        }

        if db.Arrays[a[0]] == nil {
            return errors.New("array not found")
        }

        return db.Arrays[a[0]].Insert(index, a[2])
    }

    h["MGET"] = func(a []string) error {
        if err := checkArgs(a, 2); err != nil {
            return err
            }

        if db.Queues[a[0]] == nil {
            return errors.New("queue not found")
        }

        value, err := db.Queues[a[0]].Get()
        if err == nil {
            fmt.Println(value)
        }

        return err
    }

    h["QREAD"] = func(a []string) error {
        if err := checkArgs(a, 1); err != nil {
            return err
        }

        if db.Queues[a[0]] == nil {
            return errors.New("queue not found")
        }

        printValues(db.Queues[a[0]].Values())
        return nil
    }

    // Бинарное дерево

    h["TINSERT"] = func(a []string) error {
        if err := checkArgs(a, 2); err != nil {
            return err
        }

        if db.Trees[a[0]] == nil {
            db.Trees[a[0]] = &FullBinaryTree{}
        }

        db.Trees[a[0]].Insert(a[1])
        return nil
    }

    h["TGET"] = func(a []string) error {
        if err := checkArgs(a, 2); err != nil {
            return err
        }

        if db.Trees[a[0]] == nil {
            return errors.New("tree not found")
        }

        fmt.Println(db.Trees[a[0]].Contains(a[1]))
        return nil
    }

    h["ISMEMBER"] = h["TGET"]

    h["TFULL"] = func(a []string) error {
        if err := checkArgs(a, 1); err != nil {
            return err
        }

        if db.Trees[a[0]] == nil {
            return errors.New("tree not found")
        }

        fmt.Println(db.Trees[a[0]].IsFull())
        return nil
    }

    h["TREAD"] = func(a []string) error {
        if err := checkArgs(a, 1); err != nil {
            return err
        }

        if db.Trees[a[0]] == nil {
            return errors.New("tree not found")
        }

        printValues(db.Trees[a[0]].Values())
        return nil
    }

    // Универсальная команда PRINT для любой структуры

    h["PRINT"] = func(a []string) error {
        if err := checkArgs(a, 1); err != nil {
            return err
        }

        name := a[0]

        switch {
        case db.Arrays[name] != nil:
            printValues(db.Arrays[name].Values())
        case db.SinglyLists[name] != nil:
            printValues(db.SinglyLists[name].Values())
        case db.DoublyLists[name] != nil:
            printValues(db.DoublyLists[name].Values())
        case db.Stacks[name] != nil:
            printValues(db.Stacks[name].Values())
        case db.Queues[name] != nil:
            printValues(db.Queues[name].Values())
        case db.Trees[name] != nil:
            printValues(db.Trees[name].Values())
        default:
            return errors.New("structure not found")
        }

        return nil
    }
}

// Разбор команды, вызов обработчика и сохранение изменений.
func (p *CommandProcessor) Execute(query string) error {
    // Разделение команды на имя и аргументы.
    parts := strings.Fields(query)

    if len(parts) == 0 {
        return errors.New("empty command")
    }

    // Поиск обработчика по имени команды.
    handler, exists := p.Handlers[parts[0]]

    if !exists {
        return fmt.Errorf("unknown command: %s", parts[0])
    }
