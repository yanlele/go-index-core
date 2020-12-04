package models

type Auth struct {
	ID       int    `gorm:"primary_key" json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func CheckAuth(username, password string) (int, error) {
	var auth Auth
	err := db.Select("id").Where(Auth{Username: username, Password: password}).First(&auth).Error
	if auth.ID > 0 {
		return auth.ID, nil
	}
	return -1, err
}

func GetAuthById(id int64) Auth {
	var auth Auth
	db.Where("id = ?", id).First(&auth)
	return auth
}
