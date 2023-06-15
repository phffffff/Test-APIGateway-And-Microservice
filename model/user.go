package userModel

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserWrite struct {
	Name     string `bson:"name" json:"name"`
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
	Status   int    `bson:"status" json:"status"`
}

type UserRead struct {
	Id       primitive.ObjectID `bson:"_id" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Username string             `bson:"username" json:"username"`
	Password string             `bson:"password" json:"password"`
	Status   int                `bson:"status" json:"status"`
}

type UserLogin struct {
	Username string `bson:"username" json:"username"`
	Password string `bson:"password" json:"password"`
}

func (u *UserRead) GetUserId() primitive.ObjectID {
	return u.Id
}
