package mongo

import (
	"context"
	"github.com/jxmexdev/go-todo-app/app/db"
	"github.com/jxmexdev/go-todo-app/app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type TaskRepository struct {
	db *mongo.Database
}

var taskRepositoryInstance *TaskRepository
var taskCollection = "users"

func NewTaskMongoRepository() *TaskRepository {
	if taskRepositoryInstance == nil {
		taskRepositoryInstance = &TaskRepository{db: db.GetMongoDbInstance()}
	}
	return taskRepositoryInstance
}

func (r *TaskRepository) Create(ctx context.Context, t *models.Task) (*models.Task, error) {
	t.CreatedAt = time.Now()
	t.ID = primitive.NewObjectIDFromTimestamp(time.Now()).Hex()
	id, err := primitive.ObjectIDFromHex(t.UserID)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": id}
	opts := options.FindOneAndUpdate().SetUpsert(false).SetReturnDocument(options.After)
	update := bson.M{"$push": bson.M{"tasks": t}}
	result := r.db.Collection(taskCollection).FindOneAndUpdate(ctx, filter, update, opts)
	if result.Err() != nil {
		return nil, result.Err()
	}
	return t, nil
}

func (r *TaskRepository) Update(ctx context.Context, t *models.Task) (*models.Task, error) {
	filter := bson.M{"tasks": bson.M{
		"$elemMatch": bson.M{
			"_id": t.ID,
		},
	}}
	update := bson.M{"$set": bson.M{"tasks.$.description": t.Description, "tasks.$.completed": t.Completed}}
	result, err := r.db.Collection(taskCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		return nil, mongo.ErrNoDocuments
	}
	return t, nil
}

func (r *TaskRepository) FindById(ctx context.Context, id interface{}) (*models.Task, error) {
	type UserWithTasks struct {
		Tasks []models.Task `bson:"tasks"`
	}
	var userWithTasks UserWithTasks
	//var user models.User
	filter := bson.M{"tasks": bson.M{"$elemMatch": bson.M{"_id": id}}}
	err := r.db.Collection(taskCollection).FindOne(ctx, filter).Decode(&userWithTasks)
	if err != nil {
		return nil, err
	}

	for _, task := range userWithTasks.Tasks {
		if task.ID == id {
			return &task, nil
		}
	}
	return nil, nil
}

func (r *TaskRepository) FindAll(ctx context.Context) ([]*models.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (r *TaskRepository) Close() error {
	//TODO implement me
	panic("implement me")
}

func (r *TaskRepository) FindAllByUserId(ctx context.Context, userId string) ([]*models.Task, error) {
	type UserWithTasks struct {
		Tasks []*models.Task `bson:"tasks"`
	}
	var user *UserWithTasks
	id, err := primitive.ObjectIDFromHex(userId)
	filter := bson.M{"_id": id, "tasks": bson.M{"$exists": true}}
	err = r.db.Collection(taskCollection).FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user.Tasks, nil
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	filter := bson.M{"tasks": bson.M{
		"$elemMatch": bson.M{
			"_id": id,
		},
	}}
	update := bson.M{"$pull": bson.M{"tasks": bson.M{"_id": id}}}
	result, err := r.db.Collection(taskCollection).UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 || result.ModifiedCount == 0 {
		return mongo.ErrNoDocuments
	}
	return nil
}
