package biz

import (
	"cdncloud/internal/biz/logic"

	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase, logic.NewUserLogic)
