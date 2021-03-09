package dao

import (
	"log"

	"github.com/go-xorm/xorm"

	"lottery/models"
)

type UserDao struct {
	engine *xorm.Engine
}

func NewUserDao(engine *xorm.Engine) *UserDao {
	return &UserDao{engine: engine}
}

func (d *UserDao) Get(id int) *models.LtUser {
	data := &models.LtUser{Id: id}

	ok, err := d.engine.Get(data)

	if ok && err == nil {
		return data
	} else {
		return nil
	}

}

func (d *UserDao) GetAll(page, size int) []models.LtUser {

	offset := (page - 1) * size
	dataList := make([]models.LtUser, 0)

	err := d.engine.Desc("id").Limit(size, offset).Find(&dataList)

	if err != nil {
		log.Println("black_user_dao.GetAll error=", err)
		return dataList
	} else {
		return dataList
	}
}

func (d *UserDao) CountAll() int64 {
	num, err := d.engine.Count(&models.LtUser{})
	if err != nil {
		return 0
	} else {
		return num
	}
}

// 软删除
func (d *UserDao) Delete(id int) error {
	data := &models.LtUser{Id: id, SysStatus: 1}
	_, err := d.engine.Id(data.Id).Update(data)
	return err
}

func (d *UserDao) Update(data *models.LtUser, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *UserDao) Insert(data *models.LtUser) error {
	_, err := d.engine.Insert(data)
	return err
}

func (d *UserDao) GetByUid(uid int) *models.LtUser {
	dataList := make([]models.LtUser, 0)
	err := d.engine.Where("uid=?", uid).
		Desc("id").
		Limit(1).
		Find(&dataList)

	if err != nil || len(dataList) <= 1 {
		return nil
	} else {
		return &dataList[0]
	}
}
