package dao

import (
	"time"
)

// User represents the user table in the database.
type User struct {
    Username      string `gorm:"primaryKey;size:20"`
    HashedPassword string `gorm:"unique;size:255"`
    Nama          string `gorm:"size:255"`
    Role          string `gorm:"size:30"`
    Divisi        string `gorm:"size:255"`
    Jabatan       string `gorm:"size:255"`
    CreatedAt     time.Time
    UpdatedAt     *time.Time
    IsDeleted     bool
}

type Transaksi struct {
    ID               int `gorm:"primaryKey"`
    Kategori         string `gorm:"size:255"`
    MetodePengadaan  string `gorm:"size:255"`
    CreatedAt        time.Time
    UpdatedAt        *time.Time
    Bsu              float64 `gorm:"type:numeric(10,2)"`
    Username         string `gorm:"foreignKey:Username;size:20"`
}