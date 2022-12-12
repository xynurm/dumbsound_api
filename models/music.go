package models

import "time"

type Music struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Title     string    `json:"title" gorm:"type: varchar(25)"`
	Year      string    `json:"year" gorm:"type: varchar(20)"`
	Thumbnail string    `json:"thumbnail" gorm:"type: varchar(255)"`
	Attache   string    `json:"attache" gorm:"type: varchar(255)"`
	ArtisID   int       `json:"artisId"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Artis     Artis     `json:"artis"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
