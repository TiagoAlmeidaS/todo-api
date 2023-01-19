package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"todo_project.com/internal/app/repository"
)

type Repositories struct {
	UriConnection string
	Database      string
}

type Collections struct {
	User *mongo.Collection
	Task *mongo.Collection
}

func getDatabase(uri, database string) (*mongo.Database, error) {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}
	return client.Database(database), nil
}

func getCollections(uri, database string) (*Collections, error) {
	db, err := getDatabase(uri, database)
	if err != nil {
		return nil, err
	}

	return &Collections{
		User: db.Collection("users"),
		Task: db.Collection("tasks"),
	}, nil
}

func (r *Repositories) GetRepositories() (*repository.AllRepositories, error) {
	collection, err := getCollections(r.UriConnection, r.Database)
	if err != nil {
		return nil, err
	}

	return &repository.AllRepositories{
		IUserRepository: &UserRepository{Collection: collection.User},
		ITaskRepository: &TaskRepository{Collection: collection.Task},
	}, nil
}
