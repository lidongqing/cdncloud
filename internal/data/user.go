package data

import (
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) facade.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) Save(ctx *context.Context, u *model.User) (userId int64, err error) {
	db := r.data.DataBase
	err = db.Create(&u).Error
	return u.Id, err
}

// 根据邮箱和密码获取用户信息
func (r *userRepo) GetUserInfoByEmailAndPasswd(ctx *context.Context, email string, passwd string) (user *model.User, err error) {
	db := r.data.DataBase
	err = db.Where("email = ?", email).Where("password = ?", passwd).Where("status = ?", model.USER_STATUS_NORMAL).Find(&user).Error
	return
}

// 根据手机号和密码获取用户信息
func (r *userRepo) GetUserInfoByMobileAndPasswd(ctx *context.Context, mobile string, passwd string) (user *model.User, err error) {
	db := r.data.DataBase
	err = db.Where("mobile = ?", mobile).Where("password = ?", passwd).Where("status = ?", model.USER_STATUS_NORMAL).Find(&user).Error
	return
}

// 根据邮箱获取用户信息
func (r *userRepo) GetUserInfoByEmail(ctx *context.Context, email string) (user *model.User, err error) {
	db := r.data.DataBase
	err = db.Where("email = ?", email).Find(&user).Error
	return
}

// 根据手机号获取用户信息
func (r *userRepo) GetUserInfoByMobile(ctx *context.Context, mobile string) (user *model.User, err error) {
	db := r.data.DataBase
	err = db.Where("mobile = ?", mobile).Find(&user).Error
	return
}
