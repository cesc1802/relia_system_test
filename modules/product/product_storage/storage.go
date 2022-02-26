package product_storage

import "gorm.io/gorm"

type productMySqlStorage struct {
	db *gorm.DB
}

func NewProductMysql(db *gorm.DB) *productMySqlStorage {
	return &productMySqlStorage{
		db: db,
	}
}
