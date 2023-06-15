package appContext

import "go.mongodb.org/mongo-driver/mongo"

type appContext struct {
	col       *mongo.Collection
	secretKey string
}

type AppContext interface {
	GetDBConnection() *mongo.Collection
	GetSecretKey() string
}

func NewAppContext(col *mongo.Collection, secretKey string) *appContext {
	return &appContext{col: col}
}

func (appCtx *appContext) GetDBConnection() *mongo.Collection {
	return appCtx.col
}

func (appCtx *appContext) GetSecretKey() string {
	return appCtx.secretKey
}
