package domain

import (
	"time"

	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

type ProductVariant struct {
	ID        int             `gorm:"primaryKey;autoIncrement" json:"id"`
	ProductID int             `gorm:"not null" json:"product_id"`
	Product   Product         `json:"product"`
	Size      string          `gorm:"type:varchar(50)" json:"size"`
	Color     string          `gorm:"type:varchar(50)" json:"color"`
	Stock     int             `gorm:"not null" json:"stock"`
	CreatedAt time.Time       `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time       `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func SeedProductVariants() []ProductVariant {
	sizes := []string{"S", "M", "L", "XL"}
	colors := []string{"Red", "Blue", "Green"}
	var variants []ProductVariant

	for productID := 1; productID <= 26; productID++ {
		for _, size := range sizes {
			for _, color := range colors {
				variants = append(variants, ProductVariant{
					ProductID: productID,
					Size:      size,
					Color:     color,
					Stock:     rand.Intn(50) + 1,
				})
			}
		}
	}

	return variants
}
