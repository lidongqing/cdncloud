package logic

import (
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
)

type UserLogic struct {
	userRepo facade.UserRepo
}

func NewUserLogic(userRepo facade.UserRepo) *UserLogic {
	return &UserLogic{
		userRepo: userRepo,
	}
}

// 用户注册
func (ul *UserLogic) Register(ctx *context.Context, email string, mobile string, passwd string, code string) (userId int64, err error) {
	// 校验验证码
	codeRes, err := ul.CheckCode(ctx, code)
	if err != nil {
		return 0, err
	}
	if !codeRes {
		return 0, errors.New("验证码错误")
	}

	userName := ""
	// 检查手机号是否可用
	if mobile != "" {
		isAvail, err := ul.CheckMobile(ctx, mobile)
		if err != nil {
			return 0, err
		}
		if !isAvail {
			return 0, errors.New("手机号已存在")
		}
		userName = mobile
	}

	// 检查邮箱是否可用
	if email != "" {
		isAvail, err := ul.CheckEmail(ctx, email)
		if err != nil {
			return 0, err
		}
		if !isAvail {
			return 0, errors.New("邮箱已存在")
		}
		userName = email
	}

	// 密码明文加密
	passwd, err = ul.EncryptPasswd(ctx, passwd)
	if err != nil {
		return 0, err
	}

	// 保存用户信息
	user := &model.User{
		Username: userName,
		Email:    email,
		Mobile:   mobile,
		Password: passwd,
		Status:   model.USER_STATUS_NORMAL,
	}
	userId, err = ul.userRepo.Save(ctx, user)
	return
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

// 检查邮箱是否可用
func (ul *UserLogic) CheckEmail(ctx *context.Context, email string) (bool, error) {
	user, err := ul.userRepo.GetUserInfoByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if user.Id > 0 {
		return false, nil
	}
	return true, nil
}

// 校验验证码
func (ul *UserLogic) CheckCode(ctx *context.Context, code string) (bool, error) {
	return true, nil
}

// 密码明文加密
func (ul *UserLogic) EncryptPasswd(ctx *context.Context, passwd string) (string, error) {
	md5Sum := md5.Sum([]byte(passwd))
	md5Str := hex.EncodeToString(md5Sum[:])
	return md5Str, nil
}
