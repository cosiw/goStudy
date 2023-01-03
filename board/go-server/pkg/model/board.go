package model

type Board struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	FirstName string `json:"first"`
	LastName  string `json:"last"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Location  string `json:"location"`
	Hobby     string `json:"hobby"`
}
