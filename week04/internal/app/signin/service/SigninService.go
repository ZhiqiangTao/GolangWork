package service

import (
	"fmt"
	"golangwork/week04/internal/app/signin/dao"
	"golangwork/week04/internal/app/signin/domain"
)

type ISignInService interface {
	DoSignIn(signIn *domain.SignIn) (bool, error)
}

type SigninService struct {
}

func NewSigninService() *SigninService {
	return &SigninService{}
}

func (s *SigninService) DoSignIn(signIn *domain.SignIn) (bool, error) {
	has, err := dao.HasSignIn(signIn.UserId, signIn.SignInConfigId)
	if err != nil {
		return false, err
	}
	if has {
		return false, fmt.Errorf("已经签到过，请勿重复签到")
	}

	return dao.DoSignIn(signIn)
}
