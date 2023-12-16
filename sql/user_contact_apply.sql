CREATE TABLE IF NOT EXISTS `user_contact_apply_%s`
(
    `user_id`       BIGINT  NOT NULL COMMENT '用户id',
    `apply_id`      BIGINT  NOT NULL COMMENT '申请id',
    `apply_user_id` BIGINT  NOT NULL COMMENT '申请人id',
    `to_user_id`    BIGINT  NOT NULL COMMENT '被申请人id',
    `relation_type` BIGINT  NOT NULL COMMENT '关系',
    `apply_status`  TINYINT NOT NULL COMMENT '申请状态',
    `update_time`   BIGINT  NOT NULL DEFAULT 0 COMMENT '创建时间 毫秒',
    `create_time`   BIGINT  NOT NULL DEFAULT 0 COMMENT '创建时间 毫秒',
    INDEX `User_Contact_Apply_IDX` (`user_id`, `apply_id`)
);