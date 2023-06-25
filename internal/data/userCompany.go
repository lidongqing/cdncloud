package data

import (
	"cdncloud/internal/biz/facade"
	"cdncloud/internal/model"
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type userCompany struct {
	data *Data
	log  *log.Helper
}

func NewUserCompanyRepo(data *Data, logger log.Logger) facade.UserCompanyRepo {
	return &userCompany{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userCompany) Save(ctx *context.Context, u *model.UserCompany) (id int64, err error) {
	db := r.data.DataBase
	err = db.Create(&u).Error
	return u.Id, err
}

// 根据用户id获取用户认证信息
func (r *userCompany) GetUserCompanyInfoById(ctx *context.Context, userId int64) (user *model.UserCompany, err error) {
	db := r.data.DataBase
	err = db.Where("user_id = ?", userId).Find(&user).Error
	return
}
