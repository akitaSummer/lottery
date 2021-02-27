package services

import (
	"lottery/dao"
	"lottery/datasource"

	//"lottery/dataSource"
	"lottery/models"
)

type CodeService interface {
	GetAll() []models.LtCode
	CountAll() int64
	Get(id int) *models.LtCode
	Delete(id int) error
	Update(data *models.LtCode, columns []string) error
	Insert(data *models.LtCode) error
	//UpdateByCode(data *models.LtCode, columns []string) error
	//NextUsingCode(giftId, codeId int) *models.LtCode
}

type codeService struct {
	dao *dao.CodeDao
}

func NewCodeService() CodeService {
	return &codeService{
		dao: dao.NewCodeDao(datasource.NewDbMaster()),
	}
}

func (this *codeService) GetAll() []models.LtCode {
	return this.dao.GetAll()
}

func (this *codeService) CountAll() int64 {
	return this.dao.CountAll()
}

func (this *codeService) Get(id int) *models.LtCode {
	return this.dao.Get(id)
}

func (this *codeService) Delete(id int) error {
	return this.dao.Delete(id)
}

func (this *codeService) Update(data *models.LtCode, columns []string) error {
	return this.dao.Update(data, columns)
}

func (this *codeService) Insert(data *models.LtCode) error {
	return this.dao.Insert(data)
}

//func (this *codeService) UpdateByCode(data *models.LtCode, columns []string) error {
//	return this.dao.UpdateByCode(data, columns)
//}
//
//func (this *codeService) NextUsingCode(giftId, codeId int) *models.LtCode {
//	return this.dao.NextUsingCode(giftId, codeId)
//}
