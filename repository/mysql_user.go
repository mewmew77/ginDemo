package repository

import (
	"ginDemo/model"
	"ginDemo/repository/impl"
)

var _ UserMysqlRepo = (*impl.UserMysqlImpl)(nil)

type UserMysqlRepo interface {
	AddUser(user *model.UserForMysql) error
	ListUser() (users []model.UserForMysql, err error)
}
