package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type MusicRepository interface {
	FindMusics() ([]models.Music, error)
	GetMusicID(ID int) (models.Music, error)
	CreateMusic(music models.Music) (models.Music, error)
	// UpdateUser(user models.User) (models.User, error)
	// DeleteUser(user models.User) (models.User, error)
}

func RepositoryMusic(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMusics() ([]models.Music, error) {
	var musics []models.Music
	err := r.db.Preload("Artis").Find(&musics).Error

	return musics, err
}

func (r *repository) GetMusicID(ID int) (models.Music, error) {
	var music models.Music
	err := r.db.Preload("Artis").Find(&music, ID).Error

	return music, err
}

func (r *repository) CreateMusic(music models.Music) (models.Music, error) {
	err := r.db.Preload("Artis").Create(&music).Error

	return music, err
}
