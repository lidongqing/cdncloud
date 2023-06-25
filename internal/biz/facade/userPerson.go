package facade

import (
	"context"

	"cdncloud/internal/model"
)

type UserPersonRepo interface {
	Save(ctx *context.Context, u *model.UserPerson) (id int64, err error)
	GetUserPersonInfoById(ctx *context.Context, userId int64) (user *model.UserPerson, err error)
}
