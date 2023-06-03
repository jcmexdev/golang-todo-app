package models

import "time"

type Task struct {
	ID          string    `json:"id" bson:"_id"`
	Description string    `json:"description" bson:"description"`
	Completed   bool      `json:"completed" bson:"completed"`
	UserID      string    `json:"user_id" bson:"user_id"`
	CreatedAt   time.Time `json:"created_at" bson:"created_at"`
}
