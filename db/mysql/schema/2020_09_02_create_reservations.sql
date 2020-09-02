create table reservations
(
    id int NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (id),
    user_id int NOT NULL,
    starts_at datetime NOT NULL COMMENT '開始日時',
    ends_at datetime NOT NULL COMMENT '終了日時',
    created_at datetime NOT NULL COMMENT '登録日時',
    updated_at datetime NOT NULL COMMENT '更新日時',
    deleted_at datetime COMMENT '削除日時'
)
DEFAULT CHARSET=utf8;