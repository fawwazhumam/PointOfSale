package model

import "time"

type UserRole struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"user_id" gorm:"not null;uniqueIndex:idx_user_role_pair"`
	RoleID    uint      `json:"role_id" gorm:"not null;uniqueIndex:idx_user_role_pair"`
	User      User      `gorm:"foreignKey:UserID"`
	Role      Role      `gorm:"foreignKey:RoleID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (UserRole) TableName() string {
	return "user_role"
}