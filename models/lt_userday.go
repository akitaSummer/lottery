package models

/**

// 用户每日次数表
CREATE TABLE `lt_userday` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `uid` int unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
  `day` int unsigned NOT NULL DEFAULT 0 COMMENT '日期',
  `num` int unsigned NOT NULL DEFAULT 0 COMMENT '次数',
  `sys_updated` int unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `sys_created` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type LtUserday struct {
	Id int
	Uid int
	Day int
	Num int
	SysUpdated int
	SysCreated int
}