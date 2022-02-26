package dbutil

import (
	"fmt"
	"gorm.io/gorm"
)

type FilterFunc func(query *gorm.DB) *gorm.DB

func WithCondition(conditions map[string]interface{}) FilterFunc {
	return func(query *gorm.DB) *gorm.DB {
		return query.Where(conditions)
	}
}

func QueryByStatus(status int) FilterFunc {
	return WithCondition(map[string]interface{}{
		"status": status,
	})
}
func IsNotAdmin() FilterFunc {
	return WithCondition(map[string]interface{}{
		"is_admin": 0,
	})
}

func QueryByVerifyLevel(level []int) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("verify_level IN ?", level)
	}
}

func QueryContaining(col string, value string) FilterFunc {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where(fmt.Sprintf("%s like ?", col), "%"+value+"%")
	}
}
