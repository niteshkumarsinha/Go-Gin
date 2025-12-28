package services

import (
	"context"
	"errors"

	"github.com/nitesh111sinha/apis/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserServiceImpl struct {
	userCollection *mongo.Collection
	ctx            context.Context
}

func NewUserService(userCollection *mongo.Collection, ctx context.Context) *UserServiceImpl {
	return &UserServiceImpl{
		userCollection: userCollection,
		ctx:            ctx,
	}
}

func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.userCollection.InsertOne(u.ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "user_name", Value: name}}

	err := u.userCollection.FindOne(u.ctx, query).Decode(&user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
	var users []*models.User
	cursor, err := u.userCollection.Find(u.ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(u.ctx)

	for cursor.Next(u.ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}
	
	return users, nil
}

func (u *UserServiceImpl) UpdateUser(user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: user.Name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_email", Value: user.Email}, bson.E{Key: "user_address", Value: user.Address}}}}
	result, err := u.userCollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("user not found")
	}
	return nil
}

func (u *UserServiceImpl) ResetUser(name *string, user *models.User) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	update := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "user_name", Value: user.Name}, bson.E{Key: "user_email", Value: user.Email}, bson.E{Key: "user_address", Value: user.Address}}}}
	result, err := u.userCollection.UpdateOne(u.ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("user not found")
	}
	return nil
}

func (u *UserServiceImpl) DeleteUser(name *string) error {
	filter := bson.D{bson.E{Key: "user_name", Value: name}}
	result, err := u.userCollection.DeleteOne(u.ctx, filter)
	if err != nil {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("user not found")
	}
	return nil
}
