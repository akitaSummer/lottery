package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"lottery/models"
)

type ResultDao struct {
	engine *xorm.Engine
}

func NewResultDao(engine *xorm.Engine) *ResultDao {
	return &ResultDao{engine: engine}
}

func (d *ResultDao) Get(id int) *models.LtResult {
	data := &models.LtResult{Id: id}

	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *ResultDao) GetAll(page, size int) []models.LtResult {
	offset := (page - 1) * size
	dataList := make([]models.LtResult, 0)
	err := d.engine.Desc("id").Limit(size, offset).Find(&dataList)
	if err != nil {
		return nil
	} else {
		return dataList
	}
}

func (d *ResultDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtResult{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) CountByGift(giftId int) int64 {
	num, err := d.engine.
		Where("gift_di=?", giftId).
		Count(&models.LtResult{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

func (d *ResultDao) CountByUser(uid int) int64 {
	num, err := d.engine.
		Where("uid=?", uid).
		Count(&models.LtResult{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (d *ResultDao) Delete(id int) error {
	data := &models.LtResult{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *ResultDao) Update(data *models.LtResult, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *ResultDao) Insert(data *models.LtResult) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *ResultDao) SearchByGift(giftId, page, size int) []models.LtResult {
	dataList := make([]models.LtResult, 0)

	err := d.engine.
		Where("gift_id=?", giftId).
		Limit(page, (page-1)*size).
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *ResultDao) SearchByUser(uid, page, size int) []models.LtResult {
	dataList := make([]models.LtResult, 0)

	err := d.engine.
		Where("uid=?", uid).
		Limit(page, (page-1)*size).
		Desc("id").
		Find(&dataList)

	if err != nil {
		log.Println("Result_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *ResultDao) Create(result *models.LtResult) error {
	_, err := d.engine.Insert(result)
	return err
}
