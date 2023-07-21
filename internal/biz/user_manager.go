package biz

import "github.com/jefferyjob/go-magic/internal/data/model"

type UserManager struct {
	// ...
}

func (um *UserManager) GetUserById(userId int) (*model.User, error) {
	// 实现获取用户的具体逻辑
	// ...
	return nil, nil
}

func (um *UserManager) CreateUser(user *model.User) error {
	// 实现创建用户的具体逻辑
	// ...
	return nil
}

func (um *UserManager) UpdateUser(user *model.User) error {
	// 实现更新用户的具体逻辑
	// ...
	return nil
}

func (um *UserManager) DeleteUser(userId int) error {
	// 实现删除用户的具体逻辑
	// ...
	return nil
}
