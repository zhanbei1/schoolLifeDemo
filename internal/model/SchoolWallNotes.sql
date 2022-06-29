CREATE TABLE `SchoolWallNotes` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `student_no` varchar(50) NOT NULL,
  `note_type` int(11) NOT NULL DEFAULT '1' COMMENT '标签样式编号',
  `note_title` varchar(100) NOT NULL COMMENT '标签标题',
  `note_tag` varchar(100) DEFAULT NULL COMMENT '信息标签',
  `notes_content` varchar(100) NOT NULL COMMENT '信息内容',
  `images_url` varchar(100) DEFAULT NULL COMMENT '图片URL列表，数组形式，最多9张',
  `watermark` tinyint(1) NOT NULL COMMENT '是否打标签',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `status` tinyint(4) NOT NULL COMMENT '0:信息有效，1:信息删除',
  PRIMARY KEY (`id`),
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='校园墙信息表'