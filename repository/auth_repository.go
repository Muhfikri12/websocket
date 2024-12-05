package repository

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project/database"
	"project/domain"
	"strconv"
)

type AuthRepository struct {
	db     *gorm.DB
	cacher database.Cacher
}

func NewAuthRepository(db *gorm.DB, cacher database.Cacher) *AuthRepository {
	return &AuthRepository{db: db, cacher: cacher}
}

func (repo AuthRepository) Authenticate(user domain.User) (string, string, bool, error) {
	var userFound bool
	if err := repo.db.Model(&domain.User{}).Select("count(*)>0").Where(user).Find(&userFound).Error; err != nil {
		return "", "", false, errors.New("invalid username and/or password")
	}

	if userFound {
		token := uuid.New().String()
		key := "user:" + strconv.Itoa(int(user.ID))
		if err := repo.cacher.Set(key, token); err != nil {
			return "", "", false, err
		}

		return key, token, true, nil
	}

	return "", "", false, nil
}

func (repo AuthRepository) Register(user *domain.User) error {
	return repo.db.Create(&user).Error
}
