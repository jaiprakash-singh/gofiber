package model

import (
	"time"
)

// Student structure
type Student struct {
	Phone     string    `bson:"phone"`
	Email     string    `bson:"email"`
	Status    string    `bson:"status"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
