package biz

import (
	"context"

	v1 "cdncloud/api/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// Greeter is a Greeter model.
type Greeter struct {
	Hello string
	Id    int64
}

type UserActionLog struct {
	Id         int64   `gorm:"column:id;type:bigint(20);primary_key;AUTO_INCREMENT"`
	Uid        int64   `gorm:"column:uid;type:bigint(20);NOT NULL"`
	Url        string  `gorm:"column:url;type:varchar(255);NOT NULL"`
	Request    string  `gorm:"column:request;type:varchar(255);NOT NULL"`
	Response   string  `gorm:"column:response;type:varchar(255);NOT NULL"`
	Ip         string  `gorm:"column:ip;type:varchar(255);NOT NULL"`
	TotalTime  float64 `gorm:"column:total_time;type:decimal(8,2);NOT NULL"`
	CreateTime int64   `gorm:"column:create_time;type:bigint(20);NOT NULL"`
}

// GreeterRepo is a Greater repo.
type GreeterRepo interface {
	Save(context.Context, *Greeter) (*Greeter, error)
	Update(context.Context, *Greeter) (*Greeter, error)
	FindByID(context.Context, int64) (*UserActionLog, error)
	ListByHello(context.Context, string) ([]*Greeter, error)
	ListAll(context.Context) ([]*Greeter, error)
}

// GreeterUsecase is a Greeter usecase.
type GreeterUsecase struct {
	repo GreeterRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateGreeter creates a Greeter, and returns the new Greeter.
func (uc *GreeterUsecase) CreateGreeter(ctx context.Context, g *Greeter) (*Greeter, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.Save(ctx, g)
}

func (uc *GreeterUsecase) FindByID(ctx context.Context, g *Greeter) (*UserActionLog, error) {
	uc.log.WithContext(ctx).Infof("CreateGreeter: %v", g.Hello)
	return uc.repo.FindByID(ctx, g.Id)
}
