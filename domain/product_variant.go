package domain

import (
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type ProductVariant struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int             `gorm:"not null" json:"product_id"`
	Size      string          `gorm:"type:varchar(50)" json:"size"`
	Color     string          `gorm:"type:varchar(50)" json:"color"`
	Stock     int             `gorm:"not null" json:"stock"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func SeedProductVariants() []ProductVariant {

	variants := []ProductVariant{
		{
			ProductID: 1,
			Size:      "S",
			Color:     "Red",
			Stock:     rand.Intn(50) + 1,
		},
		{
			ProductID: 2,
			Size:      "M",
			Color:     "Blue",
			Stock:     rand.Intn(50) + 1,
		},
		{
			ProductID: 3,
			Size:      "L",
			Color:     "Green",
			Stock:     rand.Intn(50) + 1,
		},
	}

	return variants
}
