package main

import (
    "flag"
    "fmt"
    "os"
)

// Обработка аргументов и запуск программы.
func main() {
    // Параметры файла базы данных и команды.
    filename := flag.String("file", "file_go.json", "database file")
    query := flag.String("query", "", "command to execute")
    flag.Parse()

    // Проверка наличия команды.
    if *query == "" {
        fmt.Fprintln(os.Stderr, "Query was not specified")
        os.Exit(1)
    }

    // Загрузка базы данных.
    db := NewDatabase()

    if err := db.Load(*filename); err != nil {
        fmt.Fprintln(os.Stderr, "ERROR:", err)
        os.Exit(1)
    }

    // Подготовка обработчика и выполнение команды.
    processor := NewCommandProcessor(db, *filename)

    if err := processor.Execute(*query); err != nil {
        fmt.Fprintln(os.Stderr, "ERROR:", err)
        os.Exit(1)
    }
}
