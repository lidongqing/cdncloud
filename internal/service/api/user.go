package api

import (
	user "cdncloud/api/v1/user"
	"cdncloud/internal/biz/logic"
	"context"
	"errors"
	"regexp"
)

type UserService struct {
	user.UnimplementedUserServer
	ul *logic.UserLogic
	ss *SessionService
}

func NewUserService(ul *logic.UserLogic, ss *SessionService) *UserService {
	return &UserService{
		ul: ul,
		ss: ss,
	}
}

// 手机号注册
func (s *UserService) RegisterByMobile(ctx context.Context, in *user.RegisterByMobileRequest) (*user.RegisterReply, error) {
	if in.Mobile == "" {
		return nil, errors.New("手机号不能为空")
	}
	if in.Password == "" {
		return nil, errors.New("密码不能为空")
	}
	//检查两次密码是否一致
	if in.Password != in.Repassword {
		return nil, errors.New("两次密码不一致")
	}
	if in.Code == "" {
		return nil, errors.New("验证码不能为空")
	}
	userId, err := s.ul.RegisterByMobile(&ctx, in.MobilePre, in.Mobile, in.Password, in.Code)
	return &user.RegisterReply{
		UserId: userId,
	}, err
}

// 邮箱注册
func (s *UserService) RegisterByEmail(ctx context.Context, in *user.RegisterByEmailRequest) (*user.RegisterReply, error) {
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	//检查两次密码是否一致
	if in.Password != in.Repassword {
		return nil, errors.New("两次密码不一致")
	}
	if in.Code == "" {
		return nil, errors.New("验证码不能为空")
	}
	userId, err := s.ul.RegisterByEmail(&ctx, in.Email, in.Password, in.Code)
	return &user.RegisterReply{
		UserId: userId,
	}, err
}

// 登录
func (s *UserService) Login(ctx context.Context, in *user.LoginRequest) (*user.EmptyReply, error) {
	if in.UserName == "" {
		return nil, errors.New("用户手机号或邮箱不能为空")
	}
	if in.Password == "" {
		return nil, errors.New("密码不能为空")
	}
	if in.Code == "" {
		return nil, errors.New("验证码不能为空")
	}
	// 区分是手机号还是邮箱
	mobile := ""
	email := ""
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, in.UserName)
	if match {
		email = in.UserName
	} else {
		mobile = in.UserName
	}

	_, err := s.ul.Login(&ctx, email, mobile, in.Password, in.Code)
	return &user.EmptyReply{}, err
}

// 通过手机号修改密码
func (s *UserService) ChangePasswdByMobile(ctx context.Context, in *user.ChangePasswdByMobileRequest) (*user.EmptyReply, error) {
	if in.Mobile == "" {
		return nil, errors.New("手机号不能为空")
	}
	//检查两次密码是否一致
	if in.Password != in.Repassword {
		return nil, errors.New("两次密码不一致")
	}
	if in.Code == "" {
		return nil, errors.New("验证码不能为空")
	}
	_, err := s.ul.ChangePasswdByMobile(&ctx, in.Mobile, in.Password, in.Code)
	return &user.EmptyReply{}, err
}

// 通过邮箱修改密码
func (s *UserService) ChangePasswdByEmail(ctx context.Context, in *user.ChangePasswdByEmailRequest) (*user.EmptyReply, error) {
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	//检查两次密码是否一致
	if in.Password != in.Repassword {
		return nil, errors.New("两次密码不一致")
	}
	if in.Code == "" {
		return nil, errors.New("验证码不能为空")
	}
	_, err := s.ul.ChangePasswdByEmail(&ctx, in.Email, in.Password, in.Code)
	return &user.EmptyReply{}, err
}

// 发送手机号验证码
func (s *UserService) SendMobileVerifyCode(ctx context.Context, in *user.SendMobileVerifyCodeRequest) (*user.EmptyReply, error) {
	if in.Mobile == "" {
		return nil, errors.New("手机号不能为空")
	}
	match, _ := regexp.MatchString(`^1[3-9]\d{9}$`, in.Mobile)
	if !match {
		return nil, errors.New("手机号格式不正确")
	}
	_, err := s.ul.SendMobileCode(&ctx, in.Mobile)
	return &user.EmptyReply{}, err
}

// 发送邮箱验证码
func (s *UserService) SendEmailVerifyCode(ctx context.Context, in *user.SendEmailVerifyCodeRequest) (*user.EmptyReply, error) {
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	match, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, in.Email)
	if !match {
		return nil, errors.New("邮箱格式不正确")
	}
	_, err := s.ul.SendEmailCode(&ctx, in.Email)
	return &user.EmptyReply{}, err
}

// 获取账户信息
func (s *UserService) GetAccountInfo(ctx context.Context, in *user.EmptyRequest) (*user.GetAccountInfoReply, error) {
	userInfo, err := s.ul.GetAccountInfo(&ctx, 0)
	return userInfo, err
}

// 个人认证
func (s *UserService) UserPersonAuth(ctx context.Context, in *user.UserPersonAuthRequest) (*user.EmptyReply, error) {
	if in.Name == "" {
		return nil, errors.New("姓名不能为空")
	}
	if in.Card == "" {
		return nil, errors.New("身份证号不能为空")
	}
	if in.Mobile == "" {
		return nil, errors.New("手机号不能为空")
	}
	_, err := s.ul.SaveUserPersonInfo(&ctx, in.Name, in.Card, in.Mobile, in.MobilePre)
	return &user.EmptyReply{}, err
}

// 个人认证信息
func (s *UserService) GetUserPersonAuthInfo(ctx context.Context, in *user.EmptyRequest) (*user.GetUserPersonAuthReply, error) {
	userInfo, err := s.ul.GetUserPersonInfo(&ctx)
	return userInfo, err
}

// 企业认证
func (s *UserService) UserCompanyAuth(ctx context.Context, in *user.UserCompanyAuthRequest) (*user.EmptyReply, error) {
	if in.Name == "" {
		return nil, errors.New("企业名称不能为空")
	}
	if in.Code == "" {
		return nil, errors.New("营业执照号不能为空")
	}
	if in.Phone == "" {
		return nil, errors.New("手机号不能为空")
	}
	if in.Ecard == "" {
		return nil, errors.New("法人身份证号不能为空")
	}
	if in.Epreson == "" {
		return nil, errors.New("法人姓名不能为空")
	}
	if in.Address == "" {
		return nil, errors.New("联系地址不能为空")
	}
	_, err := s.ul.SaveUserCompanyInfo(&ctx, in)
	return &user.EmptyReply{}, err
}

// 企业认证信息
func (s *UserService) GetUserCompanyAuthInfo(ctx context.Context, in *user.EmptyRequest) (*user.GetUserCompanyAuthReply, error) {
	userInfo, err := s.ul.GetUserCompanyInfo(&ctx)
	return userInfo, err
}

// 更新手机号
func (s *UserService) UpdateMobile(ctx context.Context, in *user.UpdateMobileRequest) (*user.EmptyReply, error) {
	if in.Mobile == "" {
		return nil, errors.New("手机号不能为空")
	}
	if in.Code == "" {
		return nil, errors.New("验证码不能为空")
	}
	_, err := s.ul.UpdateMobile(&ctx, in.Mobile, in.MobilePre, in.Code)
	return &user.EmptyReply{}, err
}

// 更新邮箱
func (s *UserService) UpdateEmail(ctx context.Context, in *user.UpdateEmailRequest) (*user.EmptyReply, error) {
	if in.Email == "" {
		return nil, errors.New("邮箱不能为空")
	}
	if in.Code == "" {
		return nil, errors.New("验证码不能为空")
	}
	_, err := s.ul.UpdateEmail(&ctx, in.Email, in.Code)
	return &user.EmptyReply{}, err
}

// 更新昵称
func (s *UserService) UpdateNickName(ctx context.Context, in *user.UpdateNickNameRequest) (*user.EmptyReply, error) {
	if in.NickName == "" {
		return nil, errors.New("昵称不能为空")
	}
	_, err := s.ul.UpdateNickName(&ctx, in.NickName)
	return &user.EmptyReply{}, err
}

// 推广链接
func (s *UserService) GetPromotionUrl(ctx context.Context, in *user.EmptyRequest) (*user.GetPromotionUrlReply, error) {
	url, err := s.ul.GetPromotionUrl(&ctx)
	return &user.GetPromotionUrlReply{
		Url: url,
	}, err
}
