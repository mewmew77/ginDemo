package repository

import (
	"context"
	"ginDemo/model"
	"ginDemo/repository/impl"
	"go.mongodb.org/mongo-driver/bson"
)

var _ UserMongoRepo = (*impl.UserMongoImpl)(nil)

type UserMongoRepo interface {
	ListUser(ctx context.Context) (users []model.UserForMongo, err error)
	AddUser(ctx context.Context, user *model.UserForMongo) error
	UpdateUser(ctx context.Context, filter bson.M, user *model.UserForMongo) error
	DeleteUser(ctx context.Context, filter bson.M) error
}
