package mongo

import (
	"context"
	"github.com/jxmexdev/go-todo-app/app/db"
	"github.com/jxmexdev/go-todo-app/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	db *mongo.Database
}

var userRepositoryInstance *UserRepository
var userCollection = "users"

func NewUserMongoRepository() *UserRepository {
	if userRepositoryInstance == nil {
		userRepositoryInstance = &UserRepository{db: db.GetMongoDbInstance()}
	}
	return userRepositoryInstance
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
	id, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return nil, err
	}
	updatedUser := bson.M{"$set": bson.M{
		"email": user.Email,
	}}
	_, err = r.db.Collection(userCollection).UpdateOne(ctx, bson.M{"_id": id}, updatedUser)
	return user, err
}

func (r *UserRepository) FindById(ctx context.Context, id interface{}) (*models.User, error) {
	var user *models.User
	id, err := primitive.ObjectIDFromHex(id.(string))
	if err != nil {
		return nil, err
	}
	err = r.db.Collection(userCollection).FindOne(ctx, bson.M{"_id": id}).Decode(&user)
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
