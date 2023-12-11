CREATE TABLE IF NOT EXISTS `user_contact_apply_%s`
(
    `id`           BIGINT  NOT NULL COMMENT '申请id',
    `user_id`      BIGINT  NOT NULL COMMENT '用户id',
    `contact_id`   BIGINT  NOT NULL COMMENT '联系人id',
    `apply_type`   BIGINT  NOT NULL COMMENT '关系',
    `apply_status` TINYINT NOT NULL COMMENT '申请状态',
    `update_time`  BIGINT  NOT NULL DEFAULT 0 COMMENT '创建时间 毫秒',
    `create_time`  BIGINT  NOT NULL DEFAULT 0 COMMENT '创建时间 毫秒',
    UNIQUE INDEX `User_Contact_Apply_IDX` (`id`),
    INDEX `User_Contact_Apply_UserId_ContactId_IDX` (`user_id`, `contact_id`)
);