package userStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (sql *sqlModel) Delete(ctx context.Context, objectId *primitive.ObjectID) (interface{}, error) {
	result, err := sql.col.UpdateByID(ctx, &objectId, bson.M{"$set": bson.M{"status": 0}})
	if err != nil {
		return nil, err
	}
	return result, nil

	//result, err := sql.col.DeleteOne(ctx, bson.M{"_id": &objectId})
	//if err != nil {
	//	return nil, err
	//}
	//return result, nil
}
