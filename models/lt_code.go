package models

/**

// 优惠券表
CREATE TABLE `lt_code` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `gift_id` int unsigned NOT NULL DEFAULT 0 COMMENT '奖品ID，关联lt_gift表',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '虚拟券编码',
  `sys_status` smallint unsigned NOT NULL DEFAULT 0 COMMENT '状态，0 正常，1 作废，2 已发放',
  `sys_created` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `sys_updated` int unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type LtCode struct {
	Id int
	GiftId int
	Code string
	SysStatus int
	SysCreated int
	SysUpdated int
}
