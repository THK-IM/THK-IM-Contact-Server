CREATE TABLE IF NOT EXISTS `user_contact_%s`
(
    `user_id`     BIGINT      NOT NULL COMMENT '用户id',
    `contact_id`  BIGINT      NOT NULL COMMENT '联系人id',
    `session_id`  BIGINT      COMMENT '会话id',
    `relation`    BIGINT      NOT NULL COMMENT '关系',
    `note_name`   VARCHAR(64) COMMENT '备注名',
    `update_time` BIGINT      NOT NULL DEFAULT 0 COMMENT '创建时间 毫秒',
    `create_time` BIGINT      NOT NULL DEFAULT 0 COMMENT '创建时间 毫秒',
    UNIQUE INDEX `User_Contact_IDX` (`user_id`, `contact_id`)
);