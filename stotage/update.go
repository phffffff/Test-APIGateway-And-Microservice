package userStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	userModel "user_service/model"
)

func (sql *sqlModel) Update(ctx context.Context, objectId *primitive.ObjectID, data *userModel.UserWrite) (interface{}, error) {
	result, err := sql.col.UpdateByID(ctx, &objectId, bson.M{"$set": &data})
	if err != nil {
		return nil, err
	}
	return result, nil
}
