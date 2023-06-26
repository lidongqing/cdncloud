package service

import (
	"cdncloud/internal/data"
	"cdncloud/internal/service/api"

	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(
	NewGreeterService,
	api.NewUserService,
	api.NewWorkOrderService,
	data.NewUserRepo,
	data.NewUserPersonRepo,
	data.NewUserCompanyRepo,
	data.NewWorkOrderRepo,
	data.NewWorkOrderDetailRepo,
)
