package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type User struct {
	Email     string `json:"email" bson:"email"`
	Tag       string `json:"tag" bson:"tag"`
	Subscribe bool   `json:"subscribe" bson:"subscribe"`
}

func (user *User) CollectionName() string {
	return "User"
}

func InsertUserDetail(user *User, ctx context.Context) (err error) {
	if _, err = DB.Collection(user.CollectionName()).InsertOne(ctx, user); err != nil {
		return err
	}
	return
}
func CheckUserEmail(email string, ctx context.Context) (err error, exist bool) {
	var user *User

	if err = DB.Collection(user.CollectionName()).FindOne(ctx, bson.M{"email": email}).Decode(&user); err != nil {
		return err, false
	}
	return nil, true
}
func GetAllUserEmails(allUsers *[]User, ctx context.Context) (err error) {
	var user User
	findOptions := options.Find()
	cur, err := DB.Collection(user.CollectionName()).Find(ctx, bson.M{"subscribe": true}, findOptions)
	if err != nil {
		return err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var value User
		if err := cur.Decode(&value); err != nil {
			return err
		}
		*allUsers = append(*allUsers, value)

	}
	return
}

func UnsubscribeEmailWithEmailId(email string, updateData map[string]interface{}, ctx context.Context) (err error) {
	var user User
	_, err = DB.Collection(user.CollectionName()).UpdateOne(ctx, bson.M{"email": email}, updateData)
	if err != nil {
		return err
	}
	return
}
