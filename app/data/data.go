package data

import "time"

// this is for raw query select
type GetUserQuerySelectItems struct {
	Username  string  `gorm:"column:username" json:"username"`
	Nama      string  `gorm:"column:nama" json:"nama"`
	Role      string  `gorm:"column:role" json:"role"`
	Divisi    *string `gorm:"column:divisi" json:"divisi"`
	Jabatan   *string `gorm:"column:jabatan" json:"jabatan"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at" json:"updated_at"`
	IsDeleted bool    `gorm:"column:is_deleted" json:"is_deleted"`
}

type GetUserRes struct {
	RespMsg  string     `json:"resp_msg"`
	RespData []GetUserQuerySelectItems `json:"resp_data"`
}

type GetUserResRawMap struct {
	RespMsg  string     `json:"resp_msg"`
	RespData []map[string]interface{} `json:"resp_data"`
}

type CreateUserReq struct {
	Username string  `json:"username"`
	Password string  `json:"password"`
	Nama     string  `json:"nama"`
	Role     string  `json:"role"`
	Divisi   *string `json:"divisi"`
	Jabatan  *string `json:"jabatan"`
}

type CreateUserResItems struct {
	Success bool `json:"success"`
}

type CreateUserRes struct {
	RespMsg  string             `json:"resp_msg"`
	RespData CreateUserResItems `json:"resp_data"`
}
