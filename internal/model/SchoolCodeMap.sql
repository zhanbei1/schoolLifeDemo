CREATE TABLE `SchoolCodeMap` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `school_code` int(11) NOT NULL,
  `school_name` varchar(100) NOT NULL,
  `school_address_code` int(11) DEFAULT '0',
  `school_desc` varchar(100) DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '0:信息有效，1:信息删除',
  PRIMARY KEY (`id`),
  UNIQUE KEY `inx_tx_scode` (`school_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学校编码对应表'