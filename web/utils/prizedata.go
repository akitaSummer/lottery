package utils

import (
	"log"
	"lottery/comm"
	"lottery/services"
)

func PrizeGift(id, leftNum int) bool {
	giftService := services.NewGiftService()
	rows, err := giftService.DecrLeftNum(id, 1)
	if rows < 1 || err != nil {
		log.Print("prizedata.PrizeGift giftService.DecrLeftNum error=", err, ", rows=", rows)
		return false
	}
	return true
}

func PrizeCodeDiff(id int, codeService services.CodeService) string {
	lockUid := 0 - id - 1000000000
	LockLucky(lockUid)
	defer UnlockLucky(lockUid)

	codeId := 0
	codeInfo := codeService.NextUsingCode(id, codeId)

	if codeInfo != nil && codeInfo.Id > 0 {
		codeInfo.SysStatus = 2
		codeInfo.SysUpdated = comm.NowUnix()
		codeService.Update(codeInfo, nil)
	} else {
		log.Println("prizedata.prizeCodeDiff num codeInfo, gift_id=", id)
		return ""
	}
	return codeInfo.Code
}
