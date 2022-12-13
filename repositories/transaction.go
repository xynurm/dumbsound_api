package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransaction() ([]models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransactionID(ID int) (models.Transaction, error)
	GetUserTransaction(userID int) (models.Transaction, error)
	UpdateTransactionStatus(status string, ID string) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransaction() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := r.db.Preload("User").Find(&transactions).Error

	return transactions, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Preload("User").Create(&transaction).Error

	return transaction, err
}

func (r *repository) GetTransactionID(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Find(&transaction, ID).Error

	return transaction, err
}

func (r *repository) GetUserTransaction(userID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.Preload("User").Find(&transaction, "user_id = ? and status =?", userID, "pending").Error

	return transaction, err
}

func (r *repository) UpdateTransactionStatus(status string, ID string) error {
	var transaction models.Transaction
	r.db.Preload("User").First(&transaction, ID)

	// If is different & Status is "success" decrement product quantity
	if status != transaction.Status && status == "success" {
		var user models.User
		r.db.Model(&user).Where("id = ?", transaction.UserID).Update("subscribe", "true")
		transaction.Status = status
	}

	err := r.db.Save(&transaction).Error

	return err
}
