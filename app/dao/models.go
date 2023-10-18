package dao

import (
	"time"
)

type Transaksi struct {
    ID          int       `gorm:"primaryKey"`
    Kategori    string    `gorm:"size:255"`
    MlModel     string    `gorm:"size:255"`
    CreatedAt   time.Time
    UpdatedAt   *time.Time `gorm:"autoUpdateTime:false"`
    Credit      int64     `gorm:"type:numeric(10,2)"`
    Username    string    `gorm:"size:20;unique"`
    User        *User     `gorm:"foreignKey:Username;references:Username"`
}

type User struct {
    Username      string `gorm:"primaryKey;size:20"`
    HashedPassword string `gorm:"size:255"`
    Nama          string `gorm:"size:255"`
    Role          string `gorm:"size:30"`
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
