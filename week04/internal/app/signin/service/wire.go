//go:build wireinject
// +build wireinject

package service

import (
	"github.com/google/wire"
)

func InitSignInService() *SigninService {
	wire.Build(NewSigninService)
	return &SigninService{}
}
