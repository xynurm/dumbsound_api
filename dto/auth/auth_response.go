package authdto

type LoginResponse struct {
	Email  string `gorm:"type: varchar(255)" json:"email"`
	Status string `gorm:"type: varchar" json:"status"`
	Token  string `gorm:"type: varchar(255)" json:"token"`
}

type RegisterResponse struct {
	Name   string `gorm:"type: varchar(255)" json:"-"`
	Email  string `gorm:"type: varchar(255)" json:"-"`
	Status string `gorm:"type: varchar(1)" json:"-"`
	Token  string `gorm:"type: varchar(255)" json:"token"`
}

type CheckAuthResponse struct {
	Id        int    `gorm:"type: int" json:"id"`
	Name      string `gorm:"type: varchar(255)" json:"fullName"`
	Email     string `gorm:"type: varchar(255)" json:"email"`
	Status    string `gorm:"type: varchar(1) "  json:"listAs"`
	Gender    string `gorm:"type: varchar(255)" json:"gender"`
	Address   string `gorm:"type: varchar(255)" json:"address"`
	Phone     string `gorm:"type: varchar(255)" json:"phone"`
	Subscribe string `gorm:"type: varchar(255)" json:"subscribe"`
}
