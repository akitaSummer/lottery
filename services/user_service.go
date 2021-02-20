package services

import (
	"lottery/dao"
	//"lottery/dataSource"
	"lottery/models"
)

type BlackUserService interface {
	GetAll() []models.LtUser
	CountAll() int64
	Get(id int) *models.LtUser
	Delete(id int) error
	Update(data *models.LtUser, columns []string) error
	Insert(data *models.LtUser) error
	GetByUid(uid int) *models.LtUser
}

type userService struct {
	dao *dao.UserDao
}

func NewUserService() BlackUserService {
	return &userService{
		//dao: dao.NewBlackUserDao(dataSource.NewMysqlMaster()),
	}
}

func (s *userService) GetAll() []models.LtUser {
	return s.dao.GetAll()
}

func (s *userService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *userService) Get(id int) *models.LtUser {
	return s.dao.Get(id)
}

func (s *userService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *userService) Update(data *models.LtUser, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *userService) Insert(data *models.LtUser) error {
	return s.dao.Insert(data)
}

func (s *userService) GetByUid(uid int) *models.LtUser {
	return s.dao.GetByUid(uid)
}
