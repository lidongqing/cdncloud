package facade

import (
	"context"

	userModel "cdncloud/internal/model"
)

type UserRepo interface {
	Save(ctx *context.Context, u *userModel.User) (userId int64, err error)
	GetUserInfoById(ctx *context.Context, userId int64) (user *userModel.User, err error)
	GetUserInfoByEmailAndPasswd(ctx *context.Context, email string, passwd string) (user *userModel.User, err error)
	GetUserInfoByMobileAndPasswd(ctx *context.Context, mobile string, passwd string) (user *userModel.User, err error)
	GetUserInfoByEmail(ctx *context.Context, email string) (user *userModel.User, err error)
	GetUserInfoByMobile(ctx *context.Context, mobile string) (user *userModel.User, err error)
	UpdateLoginFailCount(ctx *context.Context, userId int64, failNum int64) (success bool, err error)
}
