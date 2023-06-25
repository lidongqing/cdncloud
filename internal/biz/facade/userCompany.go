package facade

import (
	"context"

	"cdncloud/internal/model"
)

type UserCompanyRepo interface {
	Save(ctx *context.Context, u *model.UserCompany) (id int64, err error)
	GetUserCompanyInfoById(ctx *context.Context, userId int64) (user *model.UserCompany, err error)
}
