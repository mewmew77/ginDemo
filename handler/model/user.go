package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type AddUserRequest struct {
	Name        string `json:"name" binding:"required"` // 必传
	Description string `json:"description"`
}

type UpdateRequest struct {
	ID          primitive.ObjectID `json:"id" binding:"required"`
	Description string             `json:"description"`
}
