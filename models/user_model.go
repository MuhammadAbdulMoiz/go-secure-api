package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	UserName      *string            `json:"username" validate:"required,min=2,max=50"`
	Email         *string            `json:"email" validate:"email,required"`
	Password      *string            `json:"password" validate:"required,min=8"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Token         *string            `json:"token"`
	Refresh_token *string            `json:"refresh_token"`
	CreatedAt     *time.Time         `json:"created_at"`
	User_id       string             `json:"user_id"`
}
