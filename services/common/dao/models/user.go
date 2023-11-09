package models

import "time"

type User struct {
	ID         uint64    `gorm:"column:id;type:bigint(20);primary_key" json:"id"`
	Username   string    `gorm:"column:username;type:VARCHAR(128);NOT NULL" json:"username"`
	Password   string    `gorm:"column:password;type:VARCHAR(128);NOT NULL" json:"password"`
	RoleType   int       `gorm:"column:role_type;type:smallint;NOT NULL" json:"role_type"`
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;NOT NULL" json:"update_time"`
}

func (m *User) TableName() string {
	return "t_user"
}
