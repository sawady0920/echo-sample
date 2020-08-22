create table todos
(
    id int NOT NULL PRIMARY KEY,
    title varchar(50) NOT NULL,
    body varchar(8000)
)
DEFAULT CHARSET=utf8;


create table users
(
    id int NOT NULL PRIMARY KEY,
    user_name varchar(50) NOT NULL,
    email varchar(50) NOT NULL
)
DEFAULT CHARSET=utf8;
