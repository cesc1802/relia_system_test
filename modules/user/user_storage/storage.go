package user_storage

import "gorm.io/gorm"

type userMySqlStorage struct {
	db *gorm.DB
}

func NewUserMysql(db *gorm.DB) *userMySqlStorage {
	return &userMySqlStorage{
		db: db,
	}
}
