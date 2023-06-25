package sessions

import "github.com/google/wire"

var ProviderSet = wire.NewSet(NewSessionHandle)
