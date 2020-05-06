package model

import "time"

//Passage model
type Passage struct {
	ID         int       `json:"id"`
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreatedBy  string    `json:"created_by"`
	CreastedAt time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
