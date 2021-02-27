package services

import (
	//"fmt"

	//"lottery/comm"
	"lottery/dao"
	"lottery/datasource"

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
	//GetUserToday(uid int) *models.LtUserday
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

//func (s *userDayService) GetUserToday(uid int) *models.LtUserday {
//	y, m, d := comm.NowTime().Date()
//	strDay := fmt.Sprintf("%d%02d%02d", y, m, d)
//	return s.dao.Search(uid, strDay)
//}
