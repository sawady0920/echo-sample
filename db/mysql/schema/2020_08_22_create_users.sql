create table users
(
    id int NOT NULL AUTO_INCREMENT,
    PRIMARY KEY (id),
    user_name varchar(50) NOT NULL,
    email varchar(50) NOT NULL,
    created_at datetime NOT NULL COMMENT '登録日時',
    updated_at datetime NOT NULL COMMENT '更新日時'
)
DEFAULT CHARSET=utf8;
