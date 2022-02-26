package user_model

import (
	"github.com/pkg/errors"
	"relia_system/shared"
)

const (
	EntityName = "User"
)

var (
	ErrUserHasBeenBlocked = shared.NewCustomError(errors.New("user has been blocked"),
		"user has been blocked",
		"ErrUserHasBeenBlocked")
)

type User struct {
	shared.SQLModel `gorm:",inline"`
	Email           string `gorm:"email"`
	Password        string `gorm:"password"`
	RefreshTokenID  string `gorm:"refresh_token_id"`
	Salt            string `gorm:"salt"`
}

func (User) TableName() string {
	return "users"
}

func (u User) IsDeActivate() bool {
	return u.Status == shared.UserDeactivate
}

func (u User) GetUserId() int {
	return u.Id
}

func (u User) GetRefreshTokenId() string {
	return u.RefreshTokenID
}

func (u *User) Mask(isAdmin bool) {
	u.GenUID(shared.DBTypeUser, 1)
}

func NewUser(email, password, salt string) User {
	return User{
		Email:    email,
		Password: password,
		Salt:     salt,
	}
}
