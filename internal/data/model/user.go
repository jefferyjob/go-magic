package model

// User DB实体层
type User struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (m *User) TableName() string {
	return "user"
}
