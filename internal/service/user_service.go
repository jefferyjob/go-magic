package service

import (
	"fmt"
	"github.com/jefferyjob/go-magic/internal/biz"
	"net/http"
	"path"
	"strconv"
)

type UserService struct {
	UserManager *biz.UserManager
}

func (us *UserService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 解析请求路径，获取用户ID
	userIdStr := path.Base(r.URL.Path)
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		http.Error(w, "Invalid userId", http.StatusBadRequest)
		return
	}

	// 调用业务逻辑获取用户信息
	user, err := us.UserManager.GetUserById(userId)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	// 将用户信息写入响应
	fmt.Fprintf(w, "User: %s, ID: %d", user.Name, user.ID)
}
