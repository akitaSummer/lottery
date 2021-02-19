package models

/**

// 抽奖记录表
CREATE TABLE `lt_result` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `gift_id` int unsigned NOT NULL DEFAULT 0 COMMENT '奖品ID，关联lt_gift表',
  `gift_name` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `gift_type` int unsigned NOT NULL DEFAULT 0 COMMENT '奖品类型，同lt_gift.gtype',
  `uid` int unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `prize_code` int unsigned NOT NULL DEFAULT 0 COMMENT '抽奖编号（4位随机数）',
  `gift_data` varchar(255) NOT NULL DEFAULT '' COMMENT '获奖信息',
  `sys_status` smallint unsigned NOT NULL DEFAULT 0 COMMENT '状态，0 正常，1 删除，2 作弊',
  `sys_created` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '用户抽奖IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type LtResult struct {
	Id int
	GiftId int
	GiftName string
	GiftType int
	Uid int
	Username string
	PrizeCode int
	GiftData string
	SysStatus int
	SysCreated int
	SysIp string
}
