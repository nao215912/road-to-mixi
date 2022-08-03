package object

type User struct {
	ID     int    `db:"id" json:"-"`
	UserID int    `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
}
