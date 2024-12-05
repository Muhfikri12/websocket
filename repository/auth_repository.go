package repository

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"project/database"
	"project/domain"
	"time"
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
		token := generateToken(user.Email, "team-1")
		key := "user:" + user.Email
		if err := repo.cacher.Set(key, token); err != nil {
			return "", "", false, err
		}

		return key, token, true, nil
	}

	return "", "", false, nil
}

func generateToken(email string, secretKey string) string {
	// Gabungkan data user
	data := fmt.Sprintf("user:%s:%d", email, time.Now().Unix())

	// Buat hash HMAC menggunakan SHA-256
	h := hmac.New(sha256.New, []byte(secretKey))
	h.Write([]byte(data))
	signature := base64.URLEncoding.EncodeToString(h.Sum(nil))

	// Encode data ke Base64
	tokenData := base64.URLEncoding.EncodeToString([]byte(data))

	// Gabungkan data dan tanda tangan
	return fmt.Sprintf("%s.%s", tokenData, signature)
}
