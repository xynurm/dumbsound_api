package usersdto

type UserResponse struct {
	Token string `gorm:"type: varchar(255)" json:"token"`
}

type DeleteResponse struct {
	ID int `json:"id"`
}
