package impl

import (
	"context"
	"ginDemo/model"
	"ginDemo/repository/infra"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserMongoImpl struct {
	mongoClient *mongo.Client
}

func NewUserMongoRepo(db *mongo.Client) *UserMongoImpl {
	return &UserMongoImpl{mongoClient: db}
}

func (r *UserMongoImpl) ListUser(ctx context.Context) (users []model.UserForMongo, err error) {
	filter := bson.M{}
	res, err := r.mongoClient.Database("test").Collection("user").Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	err = res.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *UserMongoImpl) AddUser(ctx context.Context, user *model.UserForMongo) error {
	if _, err := r.mongoClient.Database("test").Collection("user").InsertOne(ctx, user); err != nil {
		return err
	}
	return nil
}

func (r *UserMongoImpl) UpdateUser(ctx context.Context, filter bson.M, user *model.UserForMongo) error {
	err := infra.Transaction(ctx, "test", func(ctx context.Context, db *mongo.Database) error {
		_, er := db.Collection("user").UpdateOne(ctx, filter, bson.M{"$set": bson.M{"description": user.Description}})
		if er != nil {
			return er
		}
		// .....可添加其它需要在事务中执行的操作.....
		return nil
	})
	return err
}

func (r *UserMongoImpl) DeleteUser(ctx context.Context, filter bson.M) error {
	_, err := r.mongoClient.Database("test").Collection("user").DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
