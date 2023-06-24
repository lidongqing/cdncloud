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

// 根据用户id获取用户信息
func (r *userRepo) GetUserInfoById(ctx *context.Context, userId int64) (user *model.User, err error) {
	db := r.data.DataBase
	err = db.Where("id = ?", userId).Find(&user).Error
	return
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

// 更新登录失败次数
func (r *userRepo) UpdateLoginFailCount(ctx *context.Context, userId int64, failNum int64) (success bool, err error) {
	db := r.data.DataBase
	err = db.Model(&model.User{}).Where("id = ?", userId).Update("login_times", failNum).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 更新密码
func (r *userRepo) UpdatePasswd(ctx *context.Context, userId int64, passwd string) (success bool, err error) {
	db := r.data.DataBase
	err = db.Model(&model.User{}).Where("id = ?", userId).Update("password", passwd).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 更新手机号
func (r *userRepo) UpdateMobile(ctx *context.Context, userId int64, mobile string, mobilePre string) (success bool, err error) {
	db := r.data.DataBase
	dataset := &model.User{
		Mobile:    mobile,
		MobilePre: mobilePre,
	}
	err = db.Model(&model.User{}).Where("id = ?", userId).Updates(dataset).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 更新邮箱
func (r *userRepo) UpdateEmail(ctx *context.Context, userId int64, email string) (success bool, err error) {
	db := r.data.DataBase
	err = db.Model(&model.User{}).Where("id = ?", userId).Update("email", email).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

// 更新昵称
func (r *userRepo) UpdateNickName(ctx *context.Context, userId int64, nickName string) (success bool, err error) {
	db := r.data.DataBase
	err = db.Model(&model.User{}).Where("id = ?", userId).Update("nickname", nickName).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
