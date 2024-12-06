package domain

import (
	"log"
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type ProductVariant struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int        `gorm:"not null" json:"product_id"`
	Size      string     `gorm:"type:varchar(50)" json:"size"`
	Color     string     `gorm:"type:varchar(50)" json:"color"`
	Stock     int        `gorm:"not null" json:"stock"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

func SeedProductVariants(db *gorm.DB, productID int) {

	variants := []ProductVariant{
		{
			ProductID: productID,
			Size:      "S",
			Color:     "Red",
			Stock:     rand.Intn(50) + 1,
		},
		{
			ProductID: productID,
			Size:      "M",
			Color:     "Blue",
			Stock:     rand.Intn(50) + 1,
		},
		{
			ProductID: productID,
			Size:      "L",
			Color:     "Green",
			Stock:     rand.Intn(50) + 1,
		},
	}

	for i, variant := range variants {
		if err := db.Create(&variant).Error; err != nil {
			log.Fatalf("failed to seed product variant #%d: %v", i+1, err)
		}
	}
}
