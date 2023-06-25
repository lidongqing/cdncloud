package data

import (
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type userPersonRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserPersonRepo(data *Data, logger log.Logger) facade.UserPersonRepo {
	return &userPersonRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userPersonRepo) Save(ctx *context.Context, u *model.UserPerson) (id int64, err error) {
	db := r.data.DataBase
	err = db.Create(&u).Error
	return u.Id, err
}

// 根据用户id获取用户认证信息
func (r *userPersonRepo) GetUserPersonInfoById(ctx *context.Context, userId int64) (user *model.UserPerson, err error) {
	db := r.data.DataBase
	err = db.Where("user_id = ?", userId).Find(&user).Error
	return
}
