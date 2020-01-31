package npk

import (
    "fmt"
    "database/sql"
    _ "github.com/lib/pq"
)

var db *sql.DB

func DBConnection(config Configuration) (*sql.DB, error) {
    if db == nil {
        var err error

        info := fmt.Sprintf(`host=%s port=%s user=%s password=%s
        dbname=%s sslmode=%s`,
        config.Database.Host,
        config.Database.Port,
        config.Database.Username,
        config.Database.Password,
        config.Database.Catalog,
        config.Database.SSLMode)

        db, err = sql.Open("postgres", info)
        if err != nil {
            return nil, err
        }
    }

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
