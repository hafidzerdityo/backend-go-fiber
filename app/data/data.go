package data

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author,omitempty"`
}

type User struct {
	Id             int64   `gorm:"column:id" json:"id"`
	Username       string  `gorm:"column:username" json:"username"`
	Name           string  `gorm:"column:name" json:"name"`
	Email          string  `gorm:"column:email" json:"email"`
	HashedPassword string  `gorm:"column:hashed_password" json:"hashed_password"`
	Gender         *string `gorm:"column:gender" json:"gender"`
	Birthdate      string  `gorm:"column:birthdate" json:"birthdate"`
	Age            int32   `gorm:"column:age" json:"age"`
	CreatedAt      string  `gorm:"column:created_at" json:"created_at"`
	Country        string  `gorm:"column:country" json:"country"`
}
