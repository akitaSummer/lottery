package services

import (
	"lottery/dao"
	"lottery/datasource"

	//"lottery/dataSource"
	"lottery/models"
)

type BlackipService interface {
	GetAll(page, size int) []models.LtBlackip
	CountAll() int64
	Get(id int) *models.LtBlackip
	Delete(id int) error
	Update(data *models.LtBlackip, columns []string) error
	Insert(data *models.LtBlackip) error
	GetByIp(ip string) *models.LtBlackip
}

type blackipService struct {
	dao *dao.BlackIpDao
}

func NewBlackipService() BlackipService {
	return &blackipService{
		dao: dao.NewBlackIpDao(datasource.NewDbMaster()),
	}
}

func (s *blackipService) GetAll(page, size int) []models.LtBlackip {
	return s.dao.GetAll(page, size)
}

func (s *blackipService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *blackipService) Get(id int) *models.LtBlackip {
	return s.dao.Get(id)
}

func (s *blackipService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *blackipService) Update(data *models.LtBlackip, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *blackipService) Insert(data *models.LtBlackip) error {
	return s.dao.Insert(data)
}

func (s *blackipService) GetByIp(ip string) *models.LtBlackip {
	return s.dao.GetByIp(ip)
}
