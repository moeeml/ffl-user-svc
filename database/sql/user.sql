create table `user`
(
    `id`         int unsigned not null auto_increment,
    `account`    varchar(30) not null unique,
    `password`   varchar(255) not null default '',
    `name`       varchar(60) not null default  '',
    `avatar`     varchar(255) not null default '',
    `create_time` int unsigned not null default 0,
    `update_time` int unsigned not null default 0,
    `status` varchar(20) not null default 'normal',
    primary key (`id`)
) engine = innodb default charset = utf8mb4 comment '用户表';