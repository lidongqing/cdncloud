package logic

import (
	"cdncloud/internal/biz/facade"
	"context"
)

type UserLogic struct {
	userRepo facade.UserRepo
}

func NewUserLogic(userRepo facade.UserRepo) *UserLogic {
	return &UserLogic{
		userRepo: userRepo,
	}
}

// 检查手机号是否可用
func (ul *UserLogic) CheckMobile(ctx *context.Context, mobile string) (bool, error) {
	user, err := ul.userRepo.GetUserInfoByMobile(ctx, mobile)
	if err != nil {
		return false, err
	}
	if user.Id > 0 {
		return false, nil
	}
	return true, nil
}
