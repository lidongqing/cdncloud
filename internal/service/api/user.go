package api

import (
	user "cdncloud/api/v1/user"
	"cdncloud/internal/biz/logic"
	"context"
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

func (s *UserService) CheckMobile(ctx context.Context, in *user.CheckMobileRequest) (*user.CheckMobileReply, error) {
	res, err := s.ul.CheckMobile(&ctx, in.Mobile)
	if err != nil {
		return nil, err
	}
	return &user.CheckMobileReply{
		IsAvail: res,
	}, nil
}
