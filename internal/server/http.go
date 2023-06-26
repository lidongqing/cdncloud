package server

import (
	v1 "cdncloud/api/v1"
	v1User "cdncloud/api/v1/user"
	v1WorkOrder "cdncloud/api/v1/workOrder"
	"cdncloud/internal/conf"
	"cdncloud/internal/service"
	"cdncloud/internal/service/api"
	"encoding/json"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, userService *api.UserService, workOrderService *api.WorkOrderService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.ResponseEncoder(responseEncoder),
		http.Middleware(
			recovery.Recovery(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}

	// rdCmd := redis.NewClient(&redis.Options{
	// 	Addr: "127.0.0.1:6379",
	// })
	// store, err := sessions.NewRedisStore(rdCmd, []byte("secret"))
	// if err != nil {
	// 	panic(err)
	// }
	// store.SetMaxAge(10 * 24 * 3600)
	// s := &api.SessionService{Store: store}

	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)
	v1.RegisterGreeterHTTPServer(srv, greeter)
	v1User.RegisterUserHTTPServer(srv, userService)
	v1WorkOrder.RegisterWorkOrderHTTPServer(srv, workOrderService)
	return srv
}

func responseEncoder(w http.ResponseWriter, r *http.Request, data interface{}) error {
	type Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
	res := &Response{
		Code:    200,
		Message: "success",
		Data:    data,
	}
	msRes, err := json.Marshal(res)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(msRes)
	return nil
}
