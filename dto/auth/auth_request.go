package authdto

type RegisterRequest struct {
	Name     string `gorm:"type: varchar" json:"fullName" validate:"required"`
	Email    string `gorm:"type: varchar" json:"email" validate:"required"`
	Password string `gorm:"type: varchar" json:"password" validate:"required"`
	Gender   string `gorm:"type: varchar" json:"gender"`
	Status   string `gorm:"type: varchar" json:"listAs"`
	Phone    string `json:"phone"`
	Address  string `json:"address" gorm:"type: varchar" `
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
