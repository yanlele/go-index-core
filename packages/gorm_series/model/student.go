package model

type Student struct {
	Id uint `json:"id"`
	Number uint `json:"number"`
	Name string `json:"name"`
	Gender string `gorm:"type:enum('1', '2');default:'1'"`
	Phone string `json:"phone"`
	Age int `json:"age"`
	ClassNumber string `gorm:"class_number;default:null"`
	Email string `json:"email"`
	Address string `json:"address"`
	Birthday int64 `gorm:"birthday"`
}

func (Student) TableName() string {
	return "student"
}
