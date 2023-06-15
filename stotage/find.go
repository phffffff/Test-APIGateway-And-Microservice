package userStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	userModel "user_service/model"
)

func (sql *sqlModel) FindDataWithCondition(ctx context.Context, cond map[string]interface{}) (*userModel.UserRead, error) {
	var result userModel.UserRead
	if err := sql.col.FindOne(ctx, cond).Decode(&result); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		return nil, err
	}
	return &result, nil
}
