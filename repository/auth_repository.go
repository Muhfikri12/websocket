package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"project/database"
	"project/domain"
	"time"

	"gorm.io/gorm"
)

type AuthRepository struct {
	db        *gorm.DB
	cacher    database.Cacher
	secretKey string
	db        *gorm.DB
	cacher    database.Cacher
	secretKey string
}

func NewAuthRepository(db *gorm.DB, cacher database.Cacher, secretKey string) *AuthRepository {
	return &AuthRepository{db: db, cacher: cacher, secretKey: secretKey}
func NewAuthRepository(db *gorm.DB, cacher database.Cacher, secretKey string) *AuthRepository {
	return &AuthRepository{db: db, cacher: cacher, secretKey: secretKey}
}

func (repo AuthRepository) Authenticate(user domain.User) (string, bool, error) {
	var userFound bool
	if err := repo.db.Model(&domain.User{}).Select("count(*)>0").Where(user).Find(&userFound).Error; err != nil {
		return "", false, errors.New("invalid username and/or password")
	}

	if userFound {
		tokenData, signature := generateToken(user.Email, repo.secretKey)
		if err := repo.cacher.Set(tokenData, signature); err != nil {
			return "", false, err
		}

		// Gabungkan data dan tanda tangan
		return fmt.Sprintf("%s.%s", tokenData, signature), true, nil
	}

	return "", false, nil
}

func generateToken(email string, secretKey string) (string, string) {
	// Gabungkan data user
	data := fmt.Sprintf("user:%s:%d", email, time.Now().Unix())

	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	tokenData := base64.URLEncoding.EncodeToString([]byte(data))
	return tokenData, signature
}

func (repo AuthRepository) Register(user *domain.User) error {
	return repo.db.Create(&user).Error
}
