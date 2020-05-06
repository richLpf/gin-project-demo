package model

import "time"

//User 用户表
type User struct {
	ID        int       `gorm:"primary_key" json:"id"`
	Username  string    `json:"Username"`
	Password  string    `json:"Password"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	IsDeleted uint      `json:"is_deleted"`
}
