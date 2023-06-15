package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
	tokenProvider "user_service/component/token"
)

type jwtProvider struct {
	secret string
}

func NewTokenJWTProvider(secret string) *jwtProvider {
	return &jwtProvider{secret: secret}
}

type myClaims struct {
	Payload tokenProvider.TokenPayload `json:"payload"`
	jwt.StandardClaims
}

func (j *jwtProvider) Generate(data tokenProvider.TokenPayload, expiry int) (*tokenProvider.Token, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims{
		data,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Second * time.Duration(expiry)).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	})

	myToken, err := token.SignedString([]byte(j.secret))
	if err != nil {
		return nil, err
	}

	return &tokenProvider.Token{
		Token:   myToken,
		Created: time.Now(),
		Expiry:  expiry,
	}, nil
}

func (j *jwtProvider) Validate(myToken string) (*tokenProvider.TokenPayload, error) {
	res, err := jwt.ParseWithClaims(myToken, &myClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(j.secret), nil
	})
	if err != nil {
		return nil, errors.New("error invalid token")
	}
	if !res.Valid {
		return nil, errors.New("error invalid token")
	}

	claims, ok := res.Claims.(*myClaims)
	if !ok {
		return nil, errors.New("error invalid token")
	}
	return &claims.Payload, nil
}
func (j *jwtProvider) String() string {
	return "JWT implement Provider"
}
