package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"fullName" gorm:"type: varchar(25)"`
	Email     string    `json:"email" gorm:"type: varchar(30)"`
	Password  string    `json:"-" gorm:"type: varchar(255)"`
	Status    string    `json:"listAs" form:"listAs" gorm:"type: varchar(1)"`
	Gender    string    `json:"gender" gorm:"type: varchar(25)"`
	Phone     string    `json:"phone" gorm:"type: varchar(20)"`
	Address   string    `json:"address" gorm:"type: varchar(255)"`
	Subscribe string    `json:"subscribe" gorm:"type: varchar(15)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

type UsersResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"fullName"`
	Email string `json:"email"`
	Role  string `json:"-"`
}

func (UsersResponse) TableName() string {
	return "users"
}
