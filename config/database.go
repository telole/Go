package config

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDatabase() {
    LoadConfig()

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
        GetEnv("DB_USER", "root"),
        GetEnv("DB_PASS", ""),
        GetEnv("DB_HOST", "127.0.0.1"),
        GetEnv("DB_PORT", "3306"),
        GetEnv("DB_NAME", "online_store"),
    )

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        panic("Gagal koneksi ke database: " + err.Error())
    }

    if err = DB.Ping(); err != nil {
        panic("Database tidak bisa diakses: " + err.Error())
    }

    fmt.Println("Database terhubung!")
}
