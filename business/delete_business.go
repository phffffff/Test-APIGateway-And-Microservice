package userBusiness

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"user_service/common"
	userModel "user_service/model"
)

type DeleteUserStore interface {
	Delete(ctx context.Context, objectId *primitive.ObjectID) (interface{}, error)
	FindDataWithCondition(ctx context.Context, cond map[string]interface{}) (*userModel.UserRead, error)
}
type deleteUserBiz struct {
	store DeleteUserStore
	req   common.Requester
}

func NewDeleteUserBiz(store DeleteUserStore, req common.Requester) *deleteUserBiz {
	return &deleteUserBiz{store: store, req: req}
}

func (biz *deleteUserBiz) DeleteUser(ctx context.Context, id primitive.ObjectID) (interface{}, error) {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"_id": id})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("user not exists")
	}
	if user.Status == 0 {
		return nil, errors.New("user deleted")
	}
	if user.Id != biz.req.GetUserId() {
		return nil, errors.New("no permission")
	}
	result, err := biz.store.Delete(ctx, &user.Id)
	if err != nil {
		return nil, errors.New("no delete user")
	}
	return result, nil
}
