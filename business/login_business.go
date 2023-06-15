package userBusiness

import (
	"context"
	"errors"
	tokenProvider "user_service/component/token"
	userModel "user_service/model"
)

type LoginStore interface {
	FindDataWithCondition(ctx context.Context, cond map[string]interface{}) (*userModel.UserRead, error)
}

type loginBiz struct {
	store         LoginStore
	tokenProvider tokenProvider.Provider
	expiry        int
}

func NewLoginBiz(store LoginStore, tokenProvider tokenProvider.Provider, expiry int) *loginBiz {
	return &loginBiz{store: store, tokenProvider: tokenProvider, expiry: expiry}
}

func (biz *loginBiz) Login(ctx context.Context, login *userModel.UserLogin) (*tokenProvider.Token, error) {
	user, err := biz.store.FindDataWithCondition(ctx, map[string]interface{}{"username": login.Username})
	if err != nil {
		return nil, err
	}
	if user.Status == 0 {
		return nil, errors.New("user deleted")
	}
	if user.Password != login.Password {
		return nil, errors.New("Username or Password invalid!")
	}
	payload := tokenProvider.TokenPayload{
		UserId: &user.Id,
	}

	accessToken, err := biz.tokenProvider.Generate(payload, biz.expiry)
	if err != nil {
		return nil, errors.New("err internal")
	}
	return accessToken, nil
}
