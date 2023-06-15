package userStorage

import (
	"context"
	userModel "user_service/model"
)

func (sql *sqlModel) Create(ctx context.Context, data *userModel.UserWrite) (interface{}, error) {
	result, err := sql.col.InsertOne(ctx, &data)
	if err != nil {
		return nil, err
	}
	return result, nil
}
