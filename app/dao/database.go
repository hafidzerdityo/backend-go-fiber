package dao

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB(user string, 
	dbname string, 
	password string, 
	host string) (*gorm.DB, error) {
    // Replace with your PostgreSQL connection details
	dsn := "user=" + user + " dbname=" + dbname + " password=" + password + " host=" + host + " port=5432 sslmode=disable"

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}
