package dao

import (
	"context"
	"gorm.io/gorm"
)

// UsersDao 数据访问层
type UsersDao struct {
	Ctx context.Context
	*gorm.DB
}

func NewUsersDao(ctx context.Context) *UsersDao {
	return &UsersDao{
		Ctx: ctx,
		//DB:  resource.DefaultDb().WithContext(ctx),
	}
}
