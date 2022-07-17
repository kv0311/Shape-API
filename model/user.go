package model

type User struct {
	ID        int64  `db:"id, primarykey, autoincrement" json:"id"`
	UserName  string `json:"user_name"`
	Password  string `json:"password"`
	FullName  string `json:"full_name"`
	Address   string `json:"address"`
	Email     string `json:"email"`
	CreatedOn int64  `json:"created_on"`
	LastLogin int64  `json:"last_login"`
}
