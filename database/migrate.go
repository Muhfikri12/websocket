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
		&domain.Admin{},
		&domain.Banner{},
	)

}

func dropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&domain.User{},
		&domain.Admin{},
		&domain.Banner{},
	)
}

func setupJoinTables(db *gorm.DB) error {
	var err error

	return err
}
