package object

//	CREATE TABLE `users` (
//		`id` bigint(20) NOT NULL AUTO_INCREMENT,
//		`user_id` int(11) NOT NULL,
//		`name` varchar(64) DEFAULT '' NOT NULL,
//	PRIMARY KEY (`id`)
//	);

type User struct {
	ID     int    `db:"id" json:"-"`
	UserID int    `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
}
