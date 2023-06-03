package mongo

import (
	"context"
	"fmt"
	"github.com/jxmexdev/go-todo-app/app/env"
	"github.com/jxmexdev/go-todo-app/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type UserRepository struct {
	db *mongo.Database
}

var userRepositoryInstance *UserRepository
var userCollection = "users"

func NewUserMongoRepository() *UserRepository {
	if userRepositoryInstance == nil {
		userRepositoryInstance.Init()
	}
	fmt.Println("Retrieving mongo repository userRepositoryInstance")
	return userRepositoryInstance
}

func (r *UserRepository) Init() {
	dbUrl := r.getMongoDBConnectionString()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(dbUrl)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println("Initialized mongo client")
	userRepositoryInstance = &UserRepository{
		db: client.Database(env.Conf.DbName),
	}
}

func (r *UserRepository) getMongoDBConnectionString() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/", env.Conf.DbUser, env.Conf.DbPassword, env.Conf.DbHost, env.Conf.DbPort)
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) (*models.User, error) {
	mongoUser := bson.M{
		"email":    user.Email,
		"password": user.Password,
	}
	res, err := r.db.Collection(userCollection).InsertOne(ctx, mongoUser)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return user, err
}

func (r *UserRepository) Update(ctx context.Context, user *models.User) (*models.User, error) {
	updatedUser := bson.M{"$set": bson.M{
		"email": user.Email,
	}}
	_, err := r.db.Collection(userCollection).UpdateOne(ctx, bson.M{"_id": user.ID}, updatedUser)
	return user, err
}

func (r *UserRepository) FindById(ctx context.Context, id interface{}) (*models.User, error) {
	var user *models.User
	err := r.db.Collection(userCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]*models.User, error) {
	var users []*models.User
	cursor, err := r.db.Collection(userCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(ctx) {
		var user *models.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) Close() error {
	//TODO implement me after find by id
	panic("implement me")
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*models.User, error) {
	var user *models.User
	err := r.db.Collection(userCollection).FindOne(ctx, bson.M{"email": email}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
