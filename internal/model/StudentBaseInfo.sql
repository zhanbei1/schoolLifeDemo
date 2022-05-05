CREATE TABLE `StudentBaseInfo` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_no` varchar(50) NOT NULL COMMENT '学号',
  `school_name` varchar(100) NOT NULL COMMENT '学校名称',
  `grade` tinyint(4) DEFAULT NULL COMMENT '年级，一年级为1，依次累加',
  `pet_name` varchar(100) DEFAULT NULL COMMENT '个人昵称',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '信息更新时间 ',
  `is_deleted` tinyint(4) DEFAULT '0' COMMENT '0:用户活跃，1:用户注销',
  `phone_no` varchar(20) DEFAULT NULL COMMENT '用户手机号',
  `icon_url` varchar(100) DEFAULT NULL COMMENT '个人头像',
  `birthday` varchar(100) DEFAULT NULL COMMENT '生日',
  `gender` tinyint(4) DEFAULT NULL COMMENT '0:女生，1:男生',
  `password` varchar(100) DEFAULT NULL COMMENT '用户登陆密码，加过密的',
  PRIMARY KEY (`id`),
  UNIQUE KEY `inx_tx_sno` (`student_no`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='学生注册基本信息表'