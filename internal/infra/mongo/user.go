package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"todo_project.com/internal/app/repository"
	"todo_project.com/internal/domain/user"
)

type UserRepository struct {
	Collection *mongo.Collection
}

type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Name     string             `bson:"name"`
	Email    string             `bson:"email"`
	Password string             `bson:"password"`
}

func domainToUser(userDomain user.User) *User {
	objectId, err := primitive.ObjectIDFromHex(userDomain.ID)
	if err != nil {
		objectId = primitive.NewObjectID()
	}

	user := User{
		ID:       objectId,
		Name:     userDomain.Name,
		Email:    userDomain.Email.String(),
		Password: userDomain.Password.String(),
	}

	return &user
}

func (u User) toDomain() *user.User {
	email, _ := user.NewEmail(u.Email)
	password, _ := user.NewPasswordHashed(u.Password)
	userDomain := user.User{
		ID:       u.ID.Hex(),
		Name:     u.Name,
		Email:    *email,
		Password: *password,
	}
	return &userDomain
}

func (r *UserRepository) Insert(domainUser user.User) (*user.User, error) {
	user := domainToUser(domainUser)
	_, err := r.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return nil, err
	}

	return user.toDomain(), nil
}

func (r *UserRepository) GetById(id string) (*user.User, error) {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, repository.ErrUserNotFound
	}

	user := User{}
	err = r.Collection.FindOne(context.Background(), bson.M{"_id": objectID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}

	return user.toDomain(), nil
}

func (r *UserRepository) GetByEmail(email string) (*user.User, error) {
	user := User{}

	err := r.Collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, repository.ErrUserNotFound
		}

		return nil, err
	}
	return user.toDomain(), nil
}
