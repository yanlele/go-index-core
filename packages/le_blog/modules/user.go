package modules

import (
	"gorm.io/gorm"
	"le-blog/bootstrap/driver"
)

type User struct {
	gorm.Model
	Name       string    `gorm:"type:varchar(25)"`
	Password   string    `gorm:"type:varchar(32); not null; default ''"`
	Salt       string    `gorm:"type:char(4); size:4; not null; default ''"`
	Email      string    `gorm:"type:varchar(100);unique_index"`
	Profession string    `gorm:"type:varchar(255); not null; default ''"`
	Avatar     string    `gorm:"type:varchar(255); not null; default ''"`
	Articles   []Article // 用户有多篇文章
}

// 更新用户信息
func (u *User) Update() error {
	return driver.DB.Save(&u).Error
}

// 更新用户密码
func (u *User) UpdatePassword() error {
	return driver.DB.Update("password", u.Password).Error
}

// 根据id获取user
func GetUserById(id int) (User, error) {
	var user User
	err := driver.DB.First(&user, id).Error
	return user, err
}
