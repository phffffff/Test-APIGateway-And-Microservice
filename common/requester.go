package common

import "go.mongodb.org/mongo-driver/bson/primitive"

type Requester interface {
	GetUserId() primitive.ObjectID
}
