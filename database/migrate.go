package database

import (
	"gorm.io/gorm"
	"project/domain"
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
		&domain.PasswordResetToken{},
		&domain.Customer{},
		&domain.Order{},
		&domain.OrderItem{},
	)

}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&domain.User{},
		&domain.PasswordResetToken{},
		&domain.Order{},
		&domain.OrderItem{},
		&domain.Customer{},
	)
}

func setupJoinTables(db *gorm.DB) error {
	var err error

	return err
}
