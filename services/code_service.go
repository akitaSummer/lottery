package services

import (
	"lottery/dao"
	"lottery/datasource"

	//"lottery/dataSource"
	"lottery/models"
)

type CodeService interface {
	GetAll(page, size int) []models.LtCode
	CountAll() int64
	CountByGift(giftId int) int64
	Search(giftId int) []models.LtCode
	Get(id int) *models.LtCode
	Delete(id int) error
	Update(code *models.LtCode, columns []string) error
	Create(code *models.LtCode) error
	NextUsingCode(giftId, codeId int) *models.LtCode
	UpdateByCode(code *models.LtCode, columns []string) error
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{dao: dao.NewCodeDao(datasource.InstanceDbMaster())}
}

func (c codeService) GetAll(page, size int) []models.LtCode {
	return c.dao.GetAll(page, size)
}

func (c codeService) CountAll() int64 {
	return c.dao.CountAll()
}

func (c codeService) CountByGift(giftId int) int64 {
	return c.dao.CountByGift(giftId)
}

func (c codeService) Search(giftId int) []models.LtCode {
	return c.dao.Search(giftId)
}

func (c codeService) Get(id int) *models.LtCode {
	return c.dao.Get(id)
}

func (c codeService) Delete(id int) error {
	return c.dao.Delete(id)
}

func (c codeService) Update(code *models.LtCode, columns []string) error {
	return c.dao.Update(code, columns)
}

func (c codeService) Create(code *models.LtCode) error {
	return c.dao.Create(code)
}

func (c *codeService) UpdateByCode(data *models.LtCode, columns []string) error {
	return c.dao.UpdateByCode(data, columns)
}

func (c *codeService) NextUsingCode(giftId, codeId int) *models.LtCode {
	return c.dao.NextUsingCode(giftId, codeId)
}
