package domain

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Image struct {
	ID        int        `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int        `gorm:"not null" json:"product_id"`
	URLPath   string     `gorm:"type:varchar(150)" json:"url_path"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}

func SeedImages(db *gorm.DB, productID int) {

	images := []Image{
		{
			ProductID: productID,
			URLPath:   "https://example.com/images/product/aneka_buah.jpg",
		},
		{
			ProductID: productID,
			URLPath:   "https://example.com/images/product/aneka_mangga.jpg",
		},
	}

	for i, image := range images {
		if err := db.Create(&image).Error; err != nil {
			log.Fatalf("failed to seed image #%d: %v", i+1, err)
		}
	}
}
