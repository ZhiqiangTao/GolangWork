package dao

import (
	"database/sql"
	"golangwork/week04/internal/app/signin/domain"
)

type dao struct {
	con *sql.DB
}

func NewDao() *dao {
	//初始化sql con
	return &dao{con: nil}
}

func HasSignIn(userId, signInConfigId int64) (bool, error) {
	// todo sql action with con
	return false, nil
}
func DoSignIn(signIn *domain.SignIn) (bool, error) {
	// todo sql action with con
	return true, nil
}
