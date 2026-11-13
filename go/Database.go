package main

import (
    "encoding/json"
    "errors"
    "os"
)

// Запись структуры для хранения в JSON.
type DiskRecord struct {
    Type   string   `json:"type"`
    Name   string   `json:"name"`
    Values []string `json:"values"`
}

// Набор структур данных с доступом по имени.
type Database struct {
    Arrays      map[string]*ArrayADT
    SinglyLists map[string]*SinglyList
    DoublyLists map[string]*DoublyList
    Stacks      map[string]*StackADT
    Queues      map[string]*QueueADT
    Trees       map[string]*FullBinaryTree
}

// Инициализация хранилищ структур.
func NewDatabase() *Database {
    return &Database{
        Arrays:      make(map[string]*ArrayADT),
        SinglyLists: make(map[string]*SinglyList),
        DoublyLists: make(map[string]*DoublyList),
        Stacks:      make(map[string]*StackADT),
        Queues:      make(map[string]*QueueADT),
        Trees:       make(map[string]*FullBinaryTree),
    }
}

// Загрузка структур из JSON-файла.
func (db *Database) Load(filename string) error {
    content, err := os.ReadFile(filename)

    // Отсутствие файла как пустая база данных.
    if errors.Is(err, os.ErrNotExist) {
        return nil
    }

    if err != nil {
        return err
    }

    var records []DiskRecord

    // Разбор содержимого JSON.
    if err := json.Unmarshal(content, &records); err != nil {
        return err
    }

    // Восстановление структур по типу и имени.
    for _, record := range records {
        for _, value := range record.Values {
            switch record.Type {
            case "M":
                if db.Arrays[record.Name] == nil {
                    db.Arrays[record.Name] = &ArrayADT{}
                }
                db.Arrays[record.Name].Push(value)

            case "F":
                if db.SinglyLists[record.Name] == nil {
                    db.SinglyLists[record.Name] = &SinglyList{}
                }
                db.SinglyLists[record.Name].PushBack(value)

            case "L":
                if db.DoublyLists[record.Name] == nil {
                    db.DoublyLists[record.Name] = &DoublyList{}
                }
                db.DoublyLists[record.Name].PushBack(value)

            case "S":
                if db.Stacks[record.Name] == nil {
                    db.Stacks[record.Name] = &StackADT{}
                }
                db.Stacks[record.Name].Push(value)

            case "Q":
                if db.Queues[record.Name] == nil {
                    db.Queues[record.Name] = &QueueADT{}
                }
                db.Queues[record.Name].Push(value)

            case "T":
                if db.Trees[record.Name] == nil {
                    db.Trees[record.Name] = &FullBinaryTree{}
                }
                db.Trees[record.Name].Insert(value)

            default:
                return errors.New("unknown structure type in database")
            }
        }
    }

    return nil
}

// Подготовка и сохранение структур в JSON-файл.
func (db *Database) Save(filename string) error {
    records := make([]DiskRecord, 0)

    // Сериализация массивов.
    for name, structure := range db.Arrays {
        records = append(records, DiskRecord{
            Type: "M", Name: name, Values: structure.Values(),
        })
    }

    // Сериализация односвязных списков.
    for name, structure := range db.SinglyLists {
        records = append(records, DiskRecord{
            Type: "F", Name: name, Values: structure.Values(),
        })
    }

    // Сериализация двусвязных списков.
    for name, structure := range db.DoublyLists {
        records = append(records, DiskRecord{
            Type: "L", Name: name, Values: structure.Values(),
        })
    }

    // Сериализация стеков.
    for name, structure := range db.Stacks {
        records = append(records, DiskRecord{
            Type: "S", Name: name, Values: structure.Values(),
        })
    }

    // Сериализация очередей.
    for name, structure := range db.Queues {
        records = append(records, DiskRecord{
            Type: "Q", Name: name, Values: structure.Values(),
        })
    }

    // Сохранение порядка добавления элементов дерева.
    for name, structure := range db.Trees {
        records = append(records, DiskRecord{
            Type: "T", Name: name,
            Values: append([]string(nil), structure.InsertionOrder...),
        })
    }

    // Формирование JSON с отступами.
    content, err := json.MarshalIndent(records, "", "    ")

    if err != nil {
        return err
    }

    // Запись файла с правами доступа 0644.
    return os.WriteFile(filename, content, 0644)
}
