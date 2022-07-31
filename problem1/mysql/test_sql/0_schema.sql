create table `users`
(
    `id`      bigint(20)             not null auto_increment,
    `user_id` int(11)                not null,
    `name`    varchar(64) default '' not null,
    primary key (`id`)
);

create table  `friend_link`
(
    `id`       bigint(20) not null auto_increment,
    `user1_id` int(11)    not null,
    `user2_id` int(11)    not null,
    check ( user1_id < user2_id ),
    unique (user1_id, user2_id),
    primary key (`id`)
);

CREATE TABLE `block_list`
(
    `id`       bigint(20) NOT NULL AUTO_INCREMENT,
    `blocking_user_id` int(11)    NOT NULL,
    `blocked_user_id` int(11)    NOT NULL,
    check ( blocking_user_id <> blocked_user_id ),
    UNIQUE (blocking_user_id, blocked_user_id),
    PRIMARY KEY (`id`)
);
