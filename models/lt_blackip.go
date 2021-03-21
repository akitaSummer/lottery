package models

/**

// IP黑名单表
CREATE TABLE `lt_blackip` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `sys_updated` int unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `sys_created` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `sys_status` smallint unsigned NOT NULL DEFAULT 0 COMMENT '状态，0 正常，1 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type LtBlackip struct {
	Id         int
	Ip         string
	BlackTime  int
	SysUpdated int
	SysCreated int
	SysStatus  int
	SysIp      string
}
