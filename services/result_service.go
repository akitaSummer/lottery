package services

import (
	"lottery/dao"
	"lottery/datasource"

	//"lottery/dataSource"
	"lottery/models"
)

type ResultService interface {
	GetAll() []models.LtResult
	CountAll() int64
	CountByGift(giftId int) int64
	CountByUser(uid int) int64
	Get(id int) *models.LtResult
	Delete(id int) error
	Update(data *models.LtResult, columns []string) error
	Insert(data *models.LtResult) error
	SearchByGift(giftId, page, size int) []models.LtResult
	SearchByUser(uid, page, size int) []models.LtResult
}

type resultService struct {
	dao *dao.ResultDao
}

func NewResultService() ResultService {
	return &resultService{
		dao: dao.NewResultDao(datasource.NewDbMaster()),
	}
}

func (s *resultService) GetAll() []models.LtResult {
	return s.dao.GetAll()
}

func (s *resultService) CountAll() int64 {
	return s.dao.CountAll()
}

func (s *resultService) CountByGift(giftId int) int64 {
	return s.dao.CountByGift(giftId)
}
func (s *resultService) CountByUser(uid int) int64 {
	return s.dao.CountByUser(uid)
}

func (s *resultService) Get(id int) *models.LtResult {
	return s.dao.Get(id)
}

func (s *resultService) Delete(id int) error {
	return s.dao.Delete(id)
}

func (s *resultService) Update(data *models.LtResult, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *resultService) Insert(data *models.LtResult) error {
	return s.dao.Insert(data)
}

func (s *resultService) SearchByGift(giftId, page, size int) []models.LtResult {
	return s.dao.SearchByGift(giftId, page, size)
}

func (s *resultService) SearchByUser(uid, page, size int) []models.LtResult {
	return s.dao.SearchByUser(uid, page, size)
}
