package logic

import (
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"math/rand"
	"time"
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

	//随机生成6位包含数字大小写字母的字符串，作为salt
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	salt := string(b)

	// 密码明文加密
	passwd, err = ul.EncryptPasswd(ctx, passwd, salt)
	if err != nil {
		return 0, err
	}

	// 保存用户信息
	user := &model.User{
		Username: userName,
		Email:    email,
		Mobile:   mobile,
		Password: passwd,
		Salt:     salt,
		Status:   model.USER_STATUS_NORMAL,
	}
	userId, err = ul.userRepo.Save(ctx, user)
	return
}

// 登录
func (ul *UserLogic) Login(ctx *context.Context, email string, mobile string, passwd string, code string) (success bool, err error) {
	if email == "" && mobile == "" {
		return false, errors.New("手机号/邮箱不能为空")
	}
	//验证码验证
	codeRes, err := ul.CheckCode(ctx, code)
	if err != nil {
		return false, err
	}
	if !codeRes {
		return false, errors.New("验证码错误")
	}

	var user *model.User
	// 根据邮箱和密码获取用户信息
	if email != "" {
		// 检查邮箱是否存在
		user, err := ul.userRepo.GetUserInfoByEmail(ctx, email)
		if err != nil {
			return false, err
		}
		if user.Id == 0 {
			//更新登录失败计数
			failNum, err := ul.UpdateLoginFailCount(ctx, user.Id, false)
			if err != nil {
				return false, err
			}
			if failNum >= 5 {
				return false, errors.New("登录失败次数过多，请稍后再试")
			}
			return false, errors.New("用户不存在")
		}
		// 密码明文加密
		passwd, err = ul.EncryptPasswd(ctx, passwd, user.Salt)
		if err != nil {
			return false, err
		}
		user, err = ul.userRepo.GetUserInfoByEmailAndPasswd(ctx, email, passwd)
		if err != nil {
			return false, err
		}
		if user.Id == 0 {
			//更新登录失败计数
			failNum, err := ul.UpdateLoginFailCount(ctx, user.Id, false)
			if err != nil {
				return false, err
			}
			if failNum >= 5 {
				return false, errors.New("登录失败次数过多，请稍后再试")
			}
			return false, errors.New("用户名或密码错误")
		}
	}

	// 根据手机号和密码获取用户信息
	if mobile != "" {
		// 检查手机号是否存在
		user, err := ul.userRepo.GetUserInfoByMobile(ctx, mobile)
		if err != nil {
			return false, err
		}
		if user.Id == 0 {
			//更新登录失败计数
			failNum, err := ul.UpdateLoginFailCount(ctx, user.Id, false)
			if err != nil {
				return false, err
			}
			if failNum >= 5 {
				return false, errors.New("登录失败次数过多，请稍后再试")
			}
			return false, errors.New("用户不存在")
		}
		// 密码明文加密
		passwd, err = ul.EncryptPasswd(ctx, passwd, user.Salt)
		if err != nil {
			return false, err
		}
		user, err = ul.userRepo.GetUserInfoByMobileAndPasswd(ctx, mobile, passwd)
		if err != nil {
			return false, err
		}
		if user.Id == 0 {
			//更新登录失败计数
			failNum, err := ul.UpdateLoginFailCount(ctx, user.Id, false)
			if err != nil {
				return false, err
			}
			if failNum >= 5 {
				return false, errors.New("登录失败次数过多，请稍后再试")
			}
			return false, errors.New("用户名或密码错误")
		}
	}
	//登录成功，更新登录失败计数
	ul.UpdateLoginFailCount(ctx, user.Id, true)
	//@todo:记录登录日志
	//@todo:更新登录时间
	//@todo:更新登录ip
	//@todo:更新登录状态
	return true, nil
}

// 更新登录失败计数，登录失败次数超过5次，将暂时禁止登录
func (ul *UserLogic) UpdateLoginFailCount(ctx *context.Context, userId int64, reset bool) (failNum int64, err error) {
	if reset {
		failNum = 0
	} else {
		user, err := ul.userRepo.GetUserInfoById(ctx, userId)
		if err != nil {
			return 0, err
		}
		failNum = user.LoginTimes + 1
	}
	_, err = ul.userRepo.UpdateLoginFailCount(ctx, userId, failNum)
	if err != nil {
		return 0, err
	}
	return failNum, nil
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
func (ul *UserLogic) EncryptPasswd(ctx *context.Context, passwd string, salt string) (string, error) {
	md5Sum := md5.Sum([]byte(passwd))
	md5Str := hex.EncodeToString(md5Sum[:])
	// 加盐
	md5Sum = md5.Sum([]byte(md5Str + salt))
	md5Str = hex.EncodeToString(md5Sum[:])
	return md5Str, nil
}

// 发送手机验证码
func (ul *UserLogic) SendMobileCode(ctx *context.Context, mobile string) (bool, error) {
	return true, nil
}

// 发送邮箱验证码
func (ul *UserLogic) SendEmailCode(ctx *context.Context, email string) (bool, error) {
	return true, nil
}
