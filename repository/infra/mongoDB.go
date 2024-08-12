// Package infra 使用单例模型来保证mongoDB的连接只初始化一次, 后续可尝试使用 Mongoose
package infra

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"sync"
	"time"
)

var (
	client *mongo.Client
	once   sync.Once
)

func GetMongoDBClient() *mongo.Client {
	once.Do(func() {
		Connect()
	})
	return client
}

func Connect() {
	// 连接mongodb
	var err error
	client, err = mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://127.0.0.1:27017").
			SetConnectTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("connect database failed, error = %v", err)
		return
	}

	// 校验是否可连通
	if err = client.Ping(context.Background(), nil); err != nil {
		fmt.Println(err)
		return
	}
	return
}

func Close() {
	if client != nil {
		err := client.Disconnect(context.Background())
		log.Fatalf("disconnext database failed, err = %v", err)
	}
}

// Transaction 执行MongoDB事务
func Transaction(ctx context.Context, dbName string, callback func(context.Context, *mongo.Database) error) error {
	if client == nil {
		return fmt.Errorf("mongo client is not initialized")
	}

	db := client.Database(dbName)
	session, err := client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	err = session.StartTransaction()
	if err != nil {
		return err
	}

	err = callback(ctx, db)
	if err != nil {
		er := session.AbortTransaction(ctx)
		if er != nil {
			return fmt.Errorf("callback failed, err = %v; abort transaction failed, err = %v", err, er)
		}
		return fmt.Errorf("callback failed, err = %v", err)
	}
	return session.CommitTransaction(ctx)
}
