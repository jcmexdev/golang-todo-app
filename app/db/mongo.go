package db

import (
	"context"
	"fmt"
	"github.com/jxmexdev/go-todo-app/app/env"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var mongoDbInstance *mongo.Database

func GetMongoDbInstance() *mongo.Database {
	if mongoDbInstance == nil {
		mongoDbInstance = initMongoDb()
	}
	return mongoDbInstance
}

func initMongoDb() *mongo.Database {
	dbUrl := getMongoDBConnectionString()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("Initialized mongo db")
	return client.Database(env.Conf.DbName)
}

func getMongoDBConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", env.Conf.DbUser, env.Conf.DbPassword, env.Conf.DbHost, env.Conf.DbPort)
}
