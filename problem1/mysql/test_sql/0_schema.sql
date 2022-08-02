create table `users`
(
    `id`      bigint(20)  not null auto_increment,
    `user_id` bigint(20)  not null,
    `name`    varchar(64) not null,
    primary key (`id`)
);

create table `follow_relation`
(
    `following_user_id` int(20) not null,
    `followed_user_id`  int(20) not null,
    check ( following_user_id <> followed_user_id ),
    primary key (following_user_id, followed_user_id)
);

CREATE TABLE `block_relation`
(
    `blocking_user_id` int(20) NOT NULL,
    `blocked_user_id`  int(20) NOT NULL,
    check ( blocking_user_id <> blocked_user_id ),
    primary key (blocking_user_id, blocked_user_id)
);
