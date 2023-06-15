package userStorage

import "go.mongodb.org/mongo-driver/mongo"

type sqlModel struct {
	col *mongo.Collection
}

func NewSqlModel(col *mongo.Collection) *sqlModel {
	return &sqlModel{col: col}
}
