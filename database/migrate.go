package database

import (
	"project/domain"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	var err error

	if err = dropTables(db); err != nil {
		return err
	}

	if err = setupJoinTables(db); err != nil {
		return err
	}

	return db.AutoMigrate(
		&domain.User{},
		&domain.Category{},
		&domain.PasswordResetToken{},
		&domain.Product{},
		&domain.ProductVariant{},
		&domain.Image{},
		&domain.Customer{},
		&domain.Order{},
		&domain.OrderItem{},
	)

}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&domain.User{},
		&domain.Category{},
		&domain.PasswordResetToken{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Customer{},
		&domain.Product{},
		&domain.ProductVariant{},
		&domain.Image{},
	)
}

func setupJoinTables(db *gorm.DB) error {
	var err error

	return err
}
