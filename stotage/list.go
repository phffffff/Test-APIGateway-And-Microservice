package userStorage

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	userModel "user_service/model"
)

func (sql *sqlModel) ListDataWithFilter(ctx context.Context, filter *userModel.Filter) ([]userModel.UserRead, error) {
	f := make([]int, len(filter.Status))
	if len(filter.Status) == 0 {
		f = append(f, 1)
	} else {
		for _, status := range filter.Status {
			f = append(f, status)
		}
	}
	cursor, err := sql.col.Find(ctx, bson.M{"status": bson.M{"$in": &f}})

	if err != nil {
		return nil, err
	}
	var list []userModel.UserRead
	if err := cursor.All(ctx, &list); err != nil {
		return nil, err
	}
	return list, nil
}
