CREATE TABLE `users`
(
    `id`      bigint(20)             NOT NULL AUTO_INCREMENT,
    `user_id` int(11)                NOT NULL,
    `name`    varchar(64) DEFAULT '' NOT NULL,
    PRIMARY KEY (`id`)
);
-- user1 user2
CREATE TABLE `friend_link`
(
    `id`       bigint(20) NOT NULL AUTO_INCREMENT,
    `user1_id` int(11)    NOT NULL,
    `user2_id` int(11)    NOT NULL,
    check ( user1_id < user2_id ),
    UNIQUE (user1_id, user2_id),
    PRIMARY KEY (`id`)
);
-- user1 user2 block
CREATE TABLE `block_list`
(
    `id`       bigint(20) NOT NULL AUTO_INCREMENT,
    `blocking_user_id` int(11)    NOT NULL,
    `blocked_user_id` int(11)    NOT NULL,
    UNIQUE (blocking_user_id, blocked_user_id),
    PRIMARY KEY (`id`)
);

