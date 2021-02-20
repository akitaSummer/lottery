package models

/**

// 用户黑名单表
CREATE TABLE `lt_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `blacktime` int unsigned NOT NULL DEFAULT 0 COMMENT '黑名单限制到期时间',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人',
  `mobile` varchar(50) NOT NULL DEFAULT '' COMMENT '手机号',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '发奖图片',
  `sys_created` int unsigned NOT NULL DEFAULT 0 COMMENT '创建时间',
  `sys_updated` int unsigned NOT NULL DEFAULT 0 COMMENT '修改时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `sys_status` smallint unsigned NOT NULL DEFAULT 0 COMMENT '状态，0 正常，1 删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

*/

type LtUser struct {
	Id         int
	Username   string
	Blacktime  int
	Realname   string
	Mobile     string
	Address    string
	SysCreated int
	SysUpdated int
	SysStatus  int
	SysIp      string
}
