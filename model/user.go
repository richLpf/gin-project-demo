package model

import "time"

//Users 用户表
type Users struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted uint      `json:"is_deleted,omitempty"`
}

//UserRoles 用户关联角色表
type UserRoles struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Namespace string    `json:"namespace"`
	User      string    `json:"user"`
	RoleID    uint      `json:"role_id"`
	Status    uint      `json:"status"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted uint      `json:"is_deleted,omitempty"`
}
