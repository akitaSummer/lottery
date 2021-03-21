package services

import (
	//"fmt"

	"fmt"
	//"lottery/comm"
	"lottery/dao"
	"lottery/datasource"
	"strconv"
	"time"

	//"lottery/dataSource"
	"lottery/models"
)

type UserDayService interface {
	GetAll() []models.LtUserday
	CountAll() int64
	Get(id int) *models.LtUserday
	Delete(id int) error
	Update(data *models.LtUserday, columns []string) error
	Insert(data *models.LtUserday) error
	GetByUid(uid int) *models.LtUserday
	GetUserToday(uid int) *models.LtUserday
	Create(user *models.LtUserday) error
}

type userDayService struct {
	dao *dao.UserDayDao
}

func NewUserDayService() UserDayService {
	return &userDayService{
		dao: dao.NewUserDayDao(datasource.NewDbMaster()),
	}
}

func (s *userDayService) GetAll() []models.LtUserday {
	return s.dao.GetAll()
}

func (s *userDayService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *userDayService) Get(id int) *models.LtUserday {
	return s.dao.Get(id)
}

func (s *userDayService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *userDayService) Update(data *models.LtUserday, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *userDayService) Insert(data *models.LtUserday) error {
	return s.dao.Insert(data)
}

func (s *userDayService) GetByUid(uid int) *models.LtUserday {
	return s.dao.GetByUid(uid)
}

func (s userDayService) GetUserToday(uid int) *models.LtUserday {
	y, m, d := time.Now().Date()
	strDate := fmt.Sprintf("%d%02d%02d", y, m, d)
	day, _ := strconv.Atoi(strDate)
	list := s.dao.Search(uid, day)
	if list != nil && len(list) > 0 {
		return &list[0]
	} else {
		return nil
	}
}

func (u userDayService) Create(user *models.LtUserday) error {
	return u.dao.Create(user)
}
