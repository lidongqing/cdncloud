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
	data.NewUserRepo,
	data.NewUserPersonRepo,
	data.NewUserCompanyRepo,
)
