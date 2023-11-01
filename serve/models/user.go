package models

type User struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Password string `json:"-"`
	Model
}

func GetUser(id int64) (*User, error) {
	user := &User{}
	err := DB.First(user, id).Error
	return user, err
}
