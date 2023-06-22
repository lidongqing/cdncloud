package api

import (
	user "cdncloud/api/v1/user"
	"cdncloud/internal/biz/logic"
	"context"
	"errors"
)

type UserService struct {
	user.UnimplementedUserServer
	ul *logic.UserLogic
}

func NewUserService(ul *logic.UserLogic) *UserService {
	return &UserService{
		ul: ul,
	}
}

// 手机号注册
func (s *UserService) RegisterByMobile(ctx context.Context, in *user.RegisterByMobileRequest) (*user.RegisterReply, error) {
	//检查两次密码是否一致
	if in.Password != in.Repassword {
		return nil, errors.New("两次密码不一致")
	}
	userId, err := s.ul.RegisterByMobile(&ctx, in.MobilePre, in.Mobile, in.Password, in.Code)
	return &user.RegisterReply{
		UserId: userId,
	}, err
}

// 邮箱注册
func (s *UserService) RegisterByEmail(ctx context.Context, in *user.RegisterByEmailRequest) (*user.RegisterReply, error) {
	//检查两次密码是否一致
	if in.Password != in.Repassword {
		return nil, errors.New("两次密码不一致")
	}
	userId, err := s.ul.RegisterByEmail(&ctx, in.Email, in.Password, in.Code)
	return &user.RegisterReply{
		UserId: userId,
	}, err
}

// 登录
func (s *UserService) Login(ctx context.Context, in *user.LoginRequest) (*user.EmptyReply, error) {
	_, err := s.ul.Login(&ctx, in.Email, in.Mobile, in.Password, in.Code)
	return &user.EmptyReply{}, err
}

// 通过手机号修改密码
func (s *UserService) ChangePasswdByMobile(ctx context.Context, in *user.ChangePasswdByMobileRequest) (*user.EmptyReply, error) {
	_, err := s.ul.ChangePasswdByMobile(&ctx, in.Mobile, in.Password, in.Code)
	return &user.EmptyReply{}, err
}

// 通过邮箱修改密码
func (s *UserService) ChangePasswdByEmail(ctx context.Context, in *user.ChangePasswdByEmailRequest) (*user.EmptyReply, error) {
	_, err := s.ul.ChangePasswdByEmail(&ctx, in.Email, in.Password, in.Code)
	return &user.EmptyReply{}, err
}

// 发送手机号验证码
func (s *UserService) SendMobileVerifyCode(ctx context.Context, in *user.SendMobileVerifyCodeRequest) (*user.EmptyReply, error) {
	_, err := s.ul.SendMobileCode(&ctx, in.Mobile)
	return &user.EmptyReply{}, err
}

// 发送邮箱验证码
func (s *UserService) SendEmailVerifyCode(ctx context.Context, in *user.SendEmailVerifyCodeRequest) (*user.EmptyReply, error) {
	_, err := s.ul.SendEmailCode(&ctx, in.Email)
	return &user.EmptyReply{}, err
}
