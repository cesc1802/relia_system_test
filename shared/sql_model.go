package shared

import "time"

type SQLModel struct {
	Id        int       `json:"-" gorm:"primaryKey;column:id"`
	PublicId  *UID      `json:"id" gorm:"-"`
	Status    int       `json:"-" gorm:"column:status;default:1;"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time `json:"-" gorm:"column:updated_at"`
}

func (m *SQLModel) GenUID(objectType int, shardId uint32) {
	uid := NewUID(uint32(m.Id), objectType, shardId)
	m.PublicId = &uid
}

type SQLModelCreate struct {
	Id       int  `json:"-" gorm:"primaryKey;column:id;"`
	PublicId *UID `json:"id" gorm:"-"`
	Status   *int `json:"status" gorm:"column:status;default:1;"`
}

func (m *SQLModelCreate) GenUID(objectType int, shardId uint32) {
	uid := NewUID(uint32(m.Id), objectType, shardId)
	m.PublicId = &uid
}
