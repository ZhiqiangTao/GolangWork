package service

import "golangwork/week02/dao"

func GetPersonById(id int64) ([]dao.User, error) {
	return dao.QueryPersonById(id)
}
