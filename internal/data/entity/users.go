package entity

// Users DB实体层
type Users struct {
	ID   int    `json:"id" gorm:"column:id"`
	Name string `json:"name" gorm:"column:name"`
}

func (m *Users) TableName() string {
	return "users"
}
