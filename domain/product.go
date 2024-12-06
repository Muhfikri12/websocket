package domain

import (
	"log"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          int        `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string     `gorm:"type:varchar(50)" json:"name"`
	SKUProduct  string     `gorm:"type:varchar(100);unique" json:"sku_product"`
	Price       float64    `gorm:"type:float" json:"price"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"index" json:"deleted_at"`

	ProductVariants []ProductVariant `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"product_variants"`
	Images          []Image          `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL" json:"images"`
}

func seedProducts(db *gorm.DB) {
	// Data dummy untuk seeding
	products := []Product{
		{
			Name:        "Product A",
			SKUProduct:  "SKU-001",
			Price:       100000,
			Description: "This is Product A description.",
		},
		{
			Name:        "Product B",
			SKUProduct:  "SKU-002",
			Price:       150000,
			Description: "This is Product B description.",
		},
		{
			Name:        "Product C",
			SKUProduct:  "SKU-003",
			Price:       200000,
			Description: "This is Product C description.",
		},
	}

	// Masukkan data ke tabel `products`
	for i, product := range products {
		if err := db.Create(&product).Error; err != nil {
			log.Fatalf("failed to seed product #%d: %v", i+1, err)
		}

		// Tambahkan varian produk
		SeedProductVariants(db, product.ID)
		SeedImages(db, product.ID)
	}
}
