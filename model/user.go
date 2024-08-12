package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserForMysql struct {
	ID          int64  `gorm:"column:id"  json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description"  json:"description"`
}

type UserForMongo struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             ` bson:"description" json:"description"`
}
