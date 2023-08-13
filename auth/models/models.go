package models

type LoginAttempt struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Id       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	RoleId   int    `db:"role_id"`
}
