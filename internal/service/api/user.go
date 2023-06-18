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

// 用户注册
func (s *UserService) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterReply, error) {
	//检查两次密码是否一致
	if in.Password != in.Repassword {
		return nil, errors.New("两次密码不一致")
	}
	userId, err := s.ul.Register(&ctx, in.Email, in.Mobile, in.Password, in.Code)
	return &user.RegisterReply{
		UserId: userId,
	}, err
}

func (s *UserService) CheckMobile(ctx context.Context, in *user.CheckMobileRequest) (*user.CheckMobileReply, error) {
	res, err := s.ul.CheckMobile(&ctx, in.Mobile)
	if err != nil {
		return nil, err
	}
	return &user.CheckMobileReply{
		IsAvail: res,
	}, nil
}
