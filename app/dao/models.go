package dao

import (
	"time"
)

type Transaksi struct {
    ID          int       `gorm:"primaryKey"`
    Kategori    string    `gorm:"size:255;not null"`
    MlModel     string    `gorm:"size:255;not null"`
    CreatedAt   time.Time
    UpdatedAt   *time.Time `gorm:"autoUpdateTime:false"`
    Credit      int64     `gorm:"type:numeric(10,2);not null"`
    Username    string    `gorm:"size:20;unique;not null"`
    User        *User     `gorm:"foreignKey:Username;references:Username"`
}

type User struct {
    Username      string `gorm:"primaryKey;size:20;not null"`
    HashedPassword string `gorm:"size:255;not null"`
    Nama          string `gorm:"size:255;not null"`
    Role          string `gorm:"size:30;not null"`
    Divisi        *string `gorm:"size:255"`
    Jabatan       *string `gorm:"size:255"`
    CreatedAt     time.Time 
    UpdatedAt     *time.Time `gorm:"autoUpdateTime:false"`
    IsDeleted     bool 
}

func (Transaksi) TableName() string {
    return "transaksi"
}

func (User) TableName() string {
    return "user"
}
