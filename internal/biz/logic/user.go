package logic

import (
	v1User "cdncloud/api/v1/user"
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/rand"
	"regexp"
	"time"
)

type UserLogic struct {
	userRepo        facade.UserRepo
	userPersonRepo  facade.UserPersonRepo
	userCompanyRepo facade.UserCompanyRepo
}

func NewUserLogic(userRepo facade.UserRepo, userPersonRepo facade.UserPersonRepo, userCompanyRepo facade.UserCompanyRepo) *UserLogic {
	return &UserLogic{
		userRepo:        userRepo,
		userPersonRepo:  userPersonRepo,
		userCompanyRepo: userCompanyRepo,
	}
}

// 手机号注册
func (ul *UserLogic) RegisterByMobile(ctx *context.Context, mobilePre string, mobile string, passwd string, code string) (userId int64, err error) {
	// 校验手机验证码
	codeRes, err := ul.CheckMobileCode(ctx, mobile, code)
	if err != nil {
		return 0, err
	}
	if !codeRes {
		return 0, errors.New("验证码错误")
	}

	// 检查手机号是否可用
	isAvail, err := ul.CheckMobile(ctx, mobile)
	if err != nil {
		return 0, err
	}
	if !isAvail {
		return 0, errors.New("手机号已存在")
	}

	// 生成密码盐
	salt, _ := ul.GenerateSalt(ctx)

	// 密码明文加密
	passwd, err = ul.EncryptPasswd(ctx, passwd, salt)
	if err != nil {
		return 0, err
	}

	// 保存用户信息
	user := &model.User{
		Username:  mobile,
		MobilePre: mobilePre,
		Mobile:    mobile,
		Password:  passwd,
		Salt:      salt,
		Status:    model.USER_STATUS_NORMAL,
	}
	userId, err = ul.userRepo.Save(ctx, user)
	return
}

// 邮箱注册
func (ul *UserLogic) RegisterByEmail(ctx *context.Context, email string, passwd string, code string) (userId int64, err error) {
	// 校验邮箱验证码
	codeRes, err := ul.CheckEmailCode(ctx, email, code)
	if err != nil {
		return 0, err
	}
	if !codeRes {
		return 0, errors.New("验证码错误")
	}

	// 检查邮箱是否可用
	isAvail, err := ul.CheckEmail(ctx, email)
	if err != nil {
		return 0, err
	}
	if !isAvail {
		return 0, errors.New("邮箱已存在")
	}

	// 生成密码盐
	salt, _ := ul.GenerateSalt(ctx)

	// 密码明文加密
	passwd, err = ul.EncryptPasswd(ctx, passwd, salt)
	if err != nil {
		return 0, err
	}

	// 保存用户信息
	user := &model.User{
		Username: email,
		Email:    email,
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

// 手机号修改密码
func (ul *UserLogic) ChangePasswdByMobile(ctx *context.Context, mobile string, passwd string, code string) (success bool, err error) {
	// 校验手机验证码
	codeRes, err := ul.CheckMobileCode(ctx, mobile, code)
	if err != nil {
		return false, err
	}
	if !codeRes {
		return false, errors.New("验证码错误")
	}

	// 检查手机号是否存在
	user, err := ul.userRepo.GetUserInfoByMobile(ctx, mobile)
	if err != nil {
		return false, err
	}
	if user.Id == 0 {
		return false, errors.New("用户不存在")
	}

	// 生成密码盐
	salt, _ := ul.GenerateSalt(ctx)

	// 密码明文加密
	passwd, err = ul.EncryptPasswd(ctx, passwd, salt)
	if err != nil {
		return false, err
	}

	// 更新密码
	_, err = ul.userRepo.UpdatePasswd(ctx, user.Id, passwd)
	if err != nil {
		return false, err
	}
	return true, nil
}

// 邮箱修改密码
func (ul *UserLogic) ChangePasswdByEmail(ctx *context.Context, email string, passwd string, code string) (success bool, err error) {
	// 校验邮箱验证码
	codeRes, err := ul.CheckEmailCode(ctx, email, code)
	if err != nil {
		return false, err
	}
	if !codeRes {
		return false, errors.New("验证码错误")
	}

	// 检查邮箱是否存在
	user, err := ul.userRepo.GetUserInfoByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	if user.Id == 0 {
		return false, errors.New("用户不存在")
	}

	// 生成密码盐
	salt, _ := ul.GenerateSalt(ctx)

	// 密码明文加密
	passwd, err = ul.EncryptPasswd(ctx, passwd, salt)
	if err != nil {
		return false, err
	}

	// 更新密码
	_, err = ul.userRepo.UpdatePasswd(ctx, user.Id, passwd)
	if err != nil {
		return false, err
	}
	return true, nil

}

// 生成密码盐
func (ul *UserLogic) GenerateSalt(ctx *context.Context) (salt string, err error) {
	//随机生成6位包含数字大小写字母的字符串，作为salt
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 6)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	salt = string(b)
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

// 生成图片验证码
func (ul *UserLogic) GenerateImageCode(ctx *context.Context) (string, error) {
	return "", nil
}

// 校验图片验证码
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

// 校验手机验证码
func (ul *UserLogic) CheckMobileCode(ctx *context.Context, mobile string, code string) (bool, error) {
	return true, nil
}

// 发送邮箱验证码
func (ul *UserLogic) SendEmailCode(ctx *context.Context, email string) (bool, error) {
	return true, nil
}

// 校验邮箱验证码
func (ul *UserLogic) CheckEmailCode(ctx *context.Context, email string, code string) (bool, error) {
	return true, nil
}

// session读取用户id
func (ul *UserLogic) GetUserIdBySession(ctx *context.Context) (userId int64, err error) {
	return 108262, nil
}

// 获取用户信息
func (ul *UserLogic) GetAccountInfo(ctx *context.Context, userId int64) (user *v1User.GetAccountInfoReply, err error) {
	if userId == 0 {
		userId, err = ul.GetUserIdBySession(ctx)
		if err != nil {
			return nil, err
		}
	}

	userData, err := ul.userRepo.GetUserInfoById(ctx, userId)
	if err != nil {
		return nil, err
	}

	return &v1User.GetAccountInfoReply{
		UserId:     userData.Id,
		UserName:   userData.Username,
		Avatar:     userData.Avatar,
		Money:      userData.Money,
		Cybermoney: userData.Cybermoney,
	}, nil

}

// 保存用户认证信息
func (ul *UserLogic) SaveUserPersonInfo(ctx *context.Context, name string, card string, mobile string, mobilePre string) (id int64, err error) {
	userId, err := ul.GetUserIdBySession(ctx)
	if err != nil {
		return 0, err
	}
	userPerson := &model.UserPerson{
		UserID:     userId,
		Name:       name,
		Card:       card,
		Mobile:     mobile,
		MobilePre:  mobilePre,
		Status:     model.USER_PERSON_STATUS_WAIT,
		Createdate: time.Now(),
	}
	return ul.userPersonRepo.Save(ctx, userPerson)
}

// 获取用户认证信息
func (ul *UserLogic) GetUserPersonInfo(ctx *context.Context) (user *v1User.GetUserPersonAuthReply, err error) {
	userId, err := ul.GetUserIdBySession(ctx)
	if err != nil {
		return nil, err
	}
	userPersonInfo, err := ul.userPersonRepo.GetUserPersonInfoById(ctx, userId)
	if err != nil {
		return nil, err
	}
	//隐藏身份证号
	if userPersonInfo.Card != "" {
		reg := regexp.MustCompile(`^(\d{4})\d+(\d{4})$`)
		userPersonInfo.Card = reg.ReplaceAllString(userPersonInfo.Card, "$1***********$2")
	}
	//隐藏手机号
	if userPersonInfo.Mobile != "" {
		reg := regexp.MustCompile(`^(\d{3})\d+(\d{4})$`)
		userPersonInfo.Mobile = reg.ReplaceAllString(userPersonInfo.Mobile, "$1****$2")
	}
	return &v1User.GetUserPersonAuthReply{
		Name:      userPersonInfo.Name,
		Card:      userPersonInfo.Card,
		Mobile:    userPersonInfo.Mobile,
		MobilePre: userPersonInfo.MobilePre,
		Status:    userPersonInfo.Status,
	}, nil
}

// 保存用户企业认证信息
func (ul *UserLogic) SaveUserCompanyInfo(ctx *context.Context, req *v1User.UserCompanyAuthRequest) (id int64, err error) {
	userId, err := ul.GetUserIdBySession(ctx)
	if err != nil {
		return 0, err
	}
	//@todo:图片上传
	imageUrl := ""
	userCompany := &model.UserCompany{
		UserID:  userId,
		Name:    req.Name,
		Code:    req.Code,
		Epreson: req.Epreson,
		Ecard:   req.Ecard,
		Phone:   req.Phone,
		Address: req.Address,
		Image:   imageUrl,
		Status:  model.USER_COMPANY_STATUS_WAIT,
	}
	return ul.userCompanyRepo.Save(ctx, userCompany)
}

// 获取用户企业认证信息
func (ul *UserLogic) GetUserCompanyInfo(ctx *context.Context) (user *v1User.GetUserCompanyAuthReply, err error) {
	userId, err := ul.GetUserIdBySession(ctx)
	if err != nil {
		return nil, err
	}
	userCompanyInfo, err := ul.userCompanyRepo.GetUserCompanyInfoById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return &v1User.GetUserCompanyAuthReply{
		Name:    userCompanyInfo.Name,
		Code:    userCompanyInfo.Code,
		Epreson: userCompanyInfo.Epreson,
		Ecard:   userCompanyInfo.Ecard,
		Phone:   userCompanyInfo.Phone,
		Address: userCompanyInfo.Address,
		Status:  userCompanyInfo.Status,
	}, nil
}

// 推广链接
func (ul *UserLogic) GetPromotionUrl(ctx *context.Context) (url string, err error) {
	userId, err := ul.GetUserIdBySession(ctx)
	if err != nil {
		return "", err
	}
	//将userId转base64encode
	code := base64.StdEncoding.EncodeToString([]byte(string(userId)))
	url = "http://cdn.cdncloud.com/index/login/register/t/" + code
	return url, nil
}
