package dao

import (
	"context"
	"gorm.io/gorm"
)

// UserDao 数据访问层
type UserDao struct {
	Ctx context.Context
	*gorm.DB
}

func NewUsersDao(ctx context.Context) *UserDao {
	return &UserDao{
		Ctx: ctx,
		//DB:  resource.DefaultDb().WithContext(ctx),
	}
}
