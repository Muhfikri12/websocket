package domain

type User struct {
	ID       uint   `gorm:"primaryKey" json:"-"`
	FullName string `json:"full_name"`
	Email    string `gorm:"unique" example:"admin@mail.com" json:"email"`
	Password string `example:"password"`
	Role     string `gorm:"default:staff" json:"role"`
}

func UserSeed() []User {
	return []User{
		{
			FullName: "Super Admin",
			Email:    "admin@mail.com",
			Password: "admin",
			Role:     "admin",
		},
	}
}
