package data

import "time"

// this is for raw query select

type GetUserQuery struct {
	Username  string  `gorm:"column:username" json:"username"`
	Nama      string  `gorm:"column:nama" json:"nama"`
	Email      string  `gorm:"column:email" json:"email"`
	Role      string  `gorm:"column:role" json:"role"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted bool    `gorm:"column:is_deleted" json:"is_deleted"`
}

type CreateUserReq struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Nama     string  `json:"nama"`
	Email     string  `json:"email"`
	Role     string  `json:"role"`
}

type CreateUserQuery struct {
	Success bool `json:"success"`
}
