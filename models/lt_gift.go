package models

/**
// 奖品表
CREATE TABLE `lt_gift` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `prize_num` int unsigned NOT NULL DEFAULT 0 COMMENT '奖品数量，0 无限量，>0 限量，<0 无奖品',
  `left_num` int unsigned NOT NULL DEFAULT 0 COMMENT '剩余奖品数量',
  `prize_code` varchar(50) NOT NULL DEFAULT '' COMMENT '0-9999表示100%， 0-0表示万分之一中奖率',
  `prize_time` int unsigned NOT NULL DEFAULT 0 COMMENT '发奖周期，D天',
  `img` varchar(255) NOT NULL DEFAULT '' COMMENT '发奖图片',
  `displayorder` int unsigned NOT NULL DEFAULT 0 COMMENT '位置序号，小的排在前面',
  `gtype` int unsigned NOT NULL DEFAULT 0 COMMENT '奖品类型，0 虚拟币，1 虚拟券，2 实物-小奖，3 实物-大奖',
  `gdata` varchar(255) NOT NULL DEFAULT '' COMMENT '扩展数据，如：虚拟币数量',
  `time_begin` int unsigned NOT NULL DEFAULT 0 COMMENT '开始时间',
  `time_end` int unsigned NOT NULL DEFAULT 0 COMMENT '结束时间',
  `prize_data` mediumtext COMMENT '发奖计划，[[时间1, 数量1], [时间2, 数量2]]',
  `prize_begin` int unsigned NOT NULL DEFAULT 0 COMMENT '发奖周期的开始',
  `prize_end` int unsigned NOT NULL DEFAULT 0 COMMENT '发奖周期的接受',
  `sys_status` smallint unsigned NOT NULL DEFAULT 0 COMMENT '状态，0 正常，1 删除',
  `sys_created` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `sys_updated` int unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人IP',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
*/

type LtGift struct {
	Id           int
	Title        string
	PrizeNum     int
	LeftNum      int
	PrizeCode    string
	PrizeTime    int
	Img          string
	Displayorder int
	Gtype        int
	Gdata        string
	TimeBegin    int
	TimeEnd      int
	PrizeData    string
	PrizeBegin   int
	PrizeEnd     int
	SysStatus    int
	SysCreated   int
	SysUpdated   int
	SysIp        string
}
