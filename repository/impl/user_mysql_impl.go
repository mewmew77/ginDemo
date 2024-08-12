package impl

import (
	"ginDemo/model"
	"gorm.io/gorm"
)

type UserMysqlImpl struct {
	mysqlDB *gorm.DB
}

func NewUserMysqlRepo(db *gorm.DB) *UserMysqlImpl {
	return &UserMysqlImpl{mysqlDB: db}
}

func (r *UserMysqlImpl) AddUser(user *model.UserForMysql) error {
	if err := r.mysqlDB.Table("user").Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserMysqlImpl) ListUser() (users []model.UserForMysql, err error) {
	if err = r.mysqlDB.Table("user").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
