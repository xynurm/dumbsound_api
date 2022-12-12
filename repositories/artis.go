package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type ArtisRepository interface {
	FindArtists() ([]models.Artis, error)
	GetArtis(ID int) (models.Artis, error)
	CreateArtis(artis models.Artis) (models.Artis, error)
	// UpdateUser(user models.User) (models.User, error)
	// DeleteUser(user models.User) (models.User, error)
}

func RepositoryArtis(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindArtists() ([]models.Artis, error) {
	var artists []models.Artis
	err := r.db.Find(&artists).Error

	return artists, err
}

func (r *repository) CreateArtis(artis models.Artis) (models.Artis, error) {
	err := r.db.Create(&artis).Error

	return artis, err
}

func (r *repository) GetArtis(ID int) (models.Artis, error) {
	var artis models.Artis
	// not yet using category relation, cause this step doesnt Belong to Many
	err := r.db.First(&artis, ID).Error

	return artis, err
}
