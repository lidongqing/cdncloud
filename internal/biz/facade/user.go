package facade

import (
	"context"

	"cdncloud/internal/model"
)

type UserRepo interface {
	Save(ctx *context.Context, u *model.User) (userId int64, err error)
	GetUserInfoById(ctx *context.Context, userId int64) (user *model.User, err error)
	GetUserInfoByEmailAndPasswd(ctx *context.Context, email string, passwd string) (user *model.User, err error)
	GetUserInfoByMobileAndPasswd(ctx *context.Context, mobile string, passwd string) (user *model.User, err error)
	GetUserInfoByEmail(ctx *context.Context, email string) (user *model.User, err error)
	GetUserInfoByMobile(ctx *context.Context, mobile string) (user *model.User, err error)
	UpdateLoginFailCount(ctx *context.Context, userId int64, failNum int64) (success bool, err error)
	UpdatePasswd(ctx *context.Context, userId int64, passwd string) (success bool, err error)
}
