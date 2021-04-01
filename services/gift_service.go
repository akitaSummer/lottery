package services

import (
	"encoding/json"
	"log"
	"lottery/comm"
	"lottery/dao"
	"lottery/datasource"
	"lottery/models"
	"strconv"
	"strings"
)

type GiftService interface {
	GetAll(useCache bool) []models.LtGift
	CountAll() int64
	Get(id int, useCache bool) *models.LtGift
	Delete(id int) error
	Update(data *models.LtGift, columns []string) error
	Create(data *models.LtGift) error
	GetAllUse(useCache bool) []models.ObjGiftPrize
	DecrLeftNum(id, num int) (int64, error)
	IncrLeftNum(id, num int) (int64, error)
}

type giftService struct {
	dao *dao.GiftDao
}

func NewGiftService() GiftService {
	return &giftService{
		dao: dao.NewGiftDao(datasource.NewDbMaster()),
	}
}

func (s *giftService) GetAll(useCache bool) []models.LtGift {
	if !useCache {
		return s.dao.GetAll()
	}
	gifts := s.getAllByCache()
	if len(gifts) < 1 {
		gifts = s.dao.GetAll()
		s.setAllByCache(gifts)
	}
	return gifts
}

func (s *giftService) CountAll() int64 {
	//return s.dao.CountAll()
	gifts := s.GetAll(true)
	return int64(len(gifts))
}

func (s *giftService) Get(id int, useCache bool) *models.LtGift {
	if !useCache {
		return s.dao.Get(id)
	}
	gifts := s.GetAll(true)
	for _, gift := range gifts {
		if gift.Id == id {
			return &gift
		}
	}
	return nil
}

func (s *giftService) Delete(id int) error {
	data := &models.LtGift{Id: id}
	s.updateByCache(data, nil)
	return s.dao.Delete(id)
}

func (s *giftService) Update(data *models.LtGift, columns []string) error {
	s.updateByCache(data, columns)
	return s.dao.Update(data, columns)
}

func (s *giftService) Create(data *models.LtGift) error {
	s.updateByCache(data, nil)
	return s.dao.Create(data)
}

func (s *giftService) GetAllUse(useCache bool) []models.ObjGiftPrize {
	list := make([]models.LtGift, 0)

	if !useCache {
		list = s.dao.GetAllUse()
	} else {
		now := comm.NowUnix()
		gifts := s.GetAll(true)
		for _, gift := range gifts {
			if gift.Id > 0 && gift.SysStatus == 0 && gift.PrizeNum >= 0 && gift.TimeBegin <= now && gift.TimeEnd >= now {
				list = append(list, gift)
			}
		}
	}

	if list != nil {
		gifts := make([]models.ObjGiftPrize, 0)
		for _, gift := range list {
			codes := strings.Split(gift.PrizeCode, "-")
			if len(codes) == 2 {
				codeA := codes[0]
				codeB := codes[1]
				a, e1 := strconv.Atoi(codeA)
				b, e2 := strconv.Atoi(codeB)
				if e1 == nil && e2 == nil && b >= a && a >= 0 && b <= 10000 {
					data := models.ObjGiftPrize{
						Id:           gift.Id,
						Title:        gift.Title,
						PrizeNum:     gift.PrizeNum,
						LeftNum:      gift.LeftNum,
						PrizeCodeA:   a,
						PrizeCodeB:   b,
						Img:          gift.Img,
						Displayorder: gift.Displayorder,
						Gtype:        gift.Gtype,
						Gdata:        gift.Gdata,
					}
					gifts = append(gifts, data)
				}
			}
		}
		return gifts
	} else {
		return []models.ObjGiftPrize{}
	}
}

func (s *giftService) DecrLeftNum(id, num int) (int64, error) {
	return s.dao.DecrLeftNum(id, num)
}

func (s *giftService) IncrLeftNum(id, num int) (int64, error) {
	return s.dao.IncrLeftNum(id, num)
}

func (s *giftService) getAllByCache() []models.LtGift {
	key := "allgift"
	rds := datasource.InstanceCache()
	rs, err := rds.Do("GET", key)
	if err != nil {
		log.Println("gift_service.getAllByCache GET key=", key, ", error=", err)
	}
	str := comm.GetString(rs, "")
	if str == "" {
		return nil
	}
	var datalist []map[string]interface{}
	err = json.Unmarshal([]byte(str), &datalist)
	if err != nil {
		log.Println("gift_service.getAllByCache json.Unmarshal error=", err)
	}
	gifts := make([]models.LtGift, len(datalist))
	for i, value := range datalist {
		data := value
		id := comm.GetInt64FromMap(data, "Id", 0)
		if id < 0 {
			gifts[i] = models.LtGift{}
		} else {
			gifts[i] = models.LtGift{
				Id:           int(id),
				Title:        comm.GetStringFromMap(data, "Title", ""),
				PrizeNum:     int(comm.GetInt64FromMap(data, "PrizeNum", 0)),
				LeftNum:      int(comm.GetInt64FromMap(data, "LeftNum", 0)),
				PrizeCode:    comm.GetStringFromMap(data, "PrizeCode", ""),
				PrizeTime:    int(comm.GetInt64FromMap(data, "PrizeTime", 0)),
				Img:          comm.GetStringFromMap(data, "Img", ""),
				Displayorder: int(comm.GetInt64FromMap(data, "Displayorder", 0)),
				Gtype:        int(comm.GetInt64FromMap(data, "Gtype", 0)),
				Gdata:        comm.GetStringFromMap(data, "Gdata", ""),
				TimeBegin:    int(comm.GetInt64FromMap(data, "TimeBegin", 0)),
				TimeEnd:      int(comm.GetInt64FromMap(data, "TimeEnd", 0)),
				//PrizeData:    "",
				PrizeBegin: int(comm.GetInt64FromMap(data, "PrizeBegin", 0)),
				PrizeEnd:   int(comm.GetInt64FromMap(data, "PrizeEnd", 0)),
				SysStatus:  int(comm.GetInt64FromMap(data, "SysStatus", 0)),
				SysCreated: int(comm.GetInt64FromMap(data, "SysCreated", 0)),
				SysUpdated: int(comm.GetInt64FromMap(data, "SysUpdated", 0)),
				SysIp:      comm.GetStringFromMap(data, "SysIp", ""),
			}
		}
	}
	return gifts
}

func (s *giftService) setAllByCache(gifts []models.LtGift) {
	strValue := ""
	if len(gifts) > 0 {
		datalist := make([]map[string]interface{}, len(gifts))
		for i, gift := range gifts {
			data := make(map[string]interface{})
			data["Id"] = gift.Id
			data["Title"] = gift.Title
			data["PrizeNum"] = gift.PrizeNum
			data["LeftNum"] = gift.LeftNum
			data["PrizeCode"] = gift.PrizeCode
			data["PrizeTime"] = gift.PrizeTime
			data["Img"] = gift.Img
			data["DisplayOrder"] = gift.Displayorder
			data["Gtype"] = gift.Gtype
			data["Gdata"] = gift.Gdata
			data["TimeBegin"] = gift.TimeBegin
			data["TimeEnd"] = gift.TimeEnd
			//data["PrizeData"] = gift.PrizeData
			data["PrizeBegin"] = gift.PrizeBegin
			data["PrizeEnd"] = gift.PrizeEnd
			data["SysStatus"] = gift.SysStatus
			data["SysCreated"] = gift.SysCreated
			data["SysUpdated"] = gift.SysUpdated
			data["SysIp"] = gift.SysIp
			datalist[i] = data
		}
		str, err := json.Marshal(datalist)
		if err != nil {
			log.Println("gift_service.setAllByCache json.Marshal error=", err)
		}
		strValue = string(str)
	}
	key := "allgift"
	rds := datasource.InstanceCache()
	_, err := rds.Do("SET", key, strValue)
	if err != nil {
		log.Println("gift_service.setAllByCache redis SET key", key, ", error= ", err)
	}
}

func (s *giftService) updateByCache(data *models.LtGift, columns []string) {
	if data == nil || data.Id <= 0 {
		return
	}
	key := "allgift"
	rds := datasource.InstanceCache()
	rds.Do("DEL", key)
}
