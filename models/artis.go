package models

import "time"

type Artis struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Name        string    `json:"name" gorm:"type: varchar(25)"`
	Old         int       `json:"old" gorm:"type: int(11)"`
	Type        string    `json:"type" gorm:"type: varchar(25)"`
	StartCareer string    `json:"startCareer" gorm:"type: varchar(20)"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
