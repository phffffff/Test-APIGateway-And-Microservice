package token

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Provider interface {
	Generate(data TokenPayload, expiry int) (*Token, error)
	Validate(token string) (*TokenPayload, error)
}

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"create"`
	Expiry  int       `json:"expiry"`
}

type TokenPayload struct {
	UserId *primitive.ObjectID `json:"user_id"`
}
