package main

import (
    "database/sql"
    "fmt"
    
    _ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "test"
    password = "test"
    dbname   = "test"
)

func main() {

    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
        "password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    err = db.Ping()
    if err != nil {
        panic(err)
    }

    fmt.Println("Successfully connected!")

    // Выполнить SQL-запрос
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
        panic(err)
    }
    
    // Обработать результат запроса
    for rows.Next() {
        var id int
        var name string
        err = rows.Scan(&id, &name)
        if err != nil {
            panic(err)
        }
        fmt.Println(id, name)
    }
}