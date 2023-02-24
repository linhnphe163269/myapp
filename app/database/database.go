package database

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
)

func DB() (*gorm.DB, error) {
    db, err := gorm.Open("postgres", "host=db port=5432 user=user dbname=db password=password sslmode=disable")
    if err != nil {
        return nil, err
    }

    db.AutoMigrate(&model.Item{})

    return db, nil
}
