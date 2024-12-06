package domain

type Customer struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}
