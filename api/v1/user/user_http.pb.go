// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1_user

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationUserChangePasswdByEmail = "/api.v1.user.User/ChangePasswdByEmail"
const OperationUserChangePasswdByMobile = "/api.v1.user.User/ChangePasswdByMobile"
const OperationUserGetEmailVerifyCode = "/api.v1.user.User/GetEmailVerifyCode"
const OperationUserGetImageVerifyCode = "/api.v1.user.User/GetImageVerifyCode"
const OperationUserLogin = "/api.v1.user.User/Login"
const OperationUserRegisterByEmail = "/api.v1.user.User/RegisterByEmail"
const OperationUserRegisterByMobile = "/api.v1.user.User/RegisterByMobile"
const OperationUserSendMobileVerifyCode = "/api.v1.user.User/SendMobileVerifyCode"

type UserHTTPServer interface {
	ChangePasswdByEmail(context.Context, *ChangePasswdByEmailRequest) (*EmptyReply, error)
	ChangePasswdByMobile(context.Context, *ChangePasswdByMobileRequest) (*EmptyReply, error)
	GetEmailVerifyCode(context.Context, *SendEmailVerifyCodeRequest) (*EmptyReply, error)
	GetImageVerifyCode(context.Context, *EmptyRequest) (*GetImageVerifyCodeReply, error)
	Login(context.Context, *LoginRequest) (*EmptyReply, error)
	RegisterByEmail(context.Context, *RegisterByEmailRequest) (*RegisterReply, error)
	RegisterByMobile(context.Context, *RegisterByMobileRequest) (*RegisterReply, error)
	SendMobileVerifyCode(context.Context, *SendMobileVerifyCodeRequest) (*EmptyReply, error)
}

func RegisterUserHTTPServer(s *http.Server, srv UserHTTPServer) {
	r := s.Route("/")
	r.POST("/api/user/registerByMobile", _User_RegisterByMobile0_HTTP_Handler(srv))
	r.POST("/api/user/registerByEmail", _User_RegisterByEmail0_HTTP_Handler(srv))
	r.POST("/api/user/login", _User_Login0_HTTP_Handler(srv))
	r.GET("/api/user/getImageVerifyCode", _User_GetImageVerifyCode0_HTTP_Handler(srv))
	r.POST("/api/user/sendMobileVerifyCode", _User_SendMobileVerifyCode0_HTTP_Handler(srv))
	r.POST("/api/user/sendEmailVerifyCode", _User_GetEmailVerifyCode0_HTTP_Handler(srv))
	r.POST("/api/user/changePasswdByMobile", _User_ChangePasswdByMobile0_HTTP_Handler(srv))
	r.POST("/api/user/changePasswdByEmail", _User_ChangePasswdByEmail0_HTTP_Handler(srv))
}

func _User_RegisterByMobile0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterByMobileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegisterByMobile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterByMobile(ctx, req.(*RegisterByMobileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _User_RegisterByEmail0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterByEmailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserRegisterByEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RegisterByEmail(ctx, req.(*RegisterByEmailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _User_Login0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in LoginRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserLogin)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Login(ctx, req.(*LoginRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetImageVerifyCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in EmptyRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetImageVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetImageVerifyCode(ctx, req.(*EmptyRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetImageVerifyCodeReply)
		return ctx.Result(200, reply)
	}
}

func _User_SendMobileVerifyCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendMobileVerifyCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserSendMobileVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SendMobileVerifyCode(ctx, req.(*SendMobileVerifyCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _User_GetEmailVerifyCode0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SendEmailVerifyCodeRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserGetEmailVerifyCode)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetEmailVerifyCode(ctx, req.(*SendEmailVerifyCodeRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _User_ChangePasswdByMobile0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangePasswdByMobileRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserChangePasswdByMobile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangePasswdByMobile(ctx, req.(*ChangePasswdByMobileRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _User_ChangePasswdByEmail0_HTTP_Handler(srv UserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangePasswdByEmailRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationUserChangePasswdByEmail)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangePasswdByEmail(ctx, req.(*ChangePasswdByEmailRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

type UserHTTPClient interface {
	ChangePasswdByEmail(ctx context.Context, req *ChangePasswdByEmailRequest, opts ...http.CallOption) (rsp *EmptyReply, err error)
	ChangePasswdByMobile(ctx context.Context, req *ChangePasswdByMobileRequest, opts ...http.CallOption) (rsp *EmptyReply, err error)
	GetEmailVerifyCode(ctx context.Context, req *SendEmailVerifyCodeRequest, opts ...http.CallOption) (rsp *EmptyReply, err error)
	GetImageVerifyCode(ctx context.Context, req *EmptyRequest, opts ...http.CallOption) (rsp *GetImageVerifyCodeReply, err error)
	Login(ctx context.Context, req *LoginRequest, opts ...http.CallOption) (rsp *EmptyReply, err error)
	RegisterByEmail(ctx context.Context, req *RegisterByEmailRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	RegisterByMobile(ctx context.Context, req *RegisterByMobileRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
	SendMobileVerifyCode(ctx context.Context, req *SendMobileVerifyCodeRequest, opts ...http.CallOption) (rsp *EmptyReply, err error)
}

type UserHTTPClientImpl struct {
	cc *http.Client
}

func NewUserHTTPClient(client *http.Client) UserHTTPClient {
	return &UserHTTPClientImpl{client}
}

func (c *UserHTTPClientImpl) ChangePasswdByEmail(ctx context.Context, in *ChangePasswdByEmailRequest, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/api/user/changePasswdByEmail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserChangePasswdByEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) ChangePasswdByMobile(ctx context.Context, in *ChangePasswdByMobileRequest, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/api/user/changePasswdByMobile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserChangePasswdByMobile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetEmailVerifyCode(ctx context.Context, in *SendEmailVerifyCodeRequest, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/api/user/sendEmailVerifyCode"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserGetEmailVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) GetImageVerifyCode(ctx context.Context, in *EmptyRequest, opts ...http.CallOption) (*GetImageVerifyCodeReply, error) {
	var out GetImageVerifyCodeReply
	pattern := "/api/user/getImageVerifyCode"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationUserGetImageVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) Login(ctx context.Context, in *LoginRequest, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/api/user/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserLogin))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RegisterByEmail(ctx context.Context, in *RegisterByEmailRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/user/registerByEmail"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegisterByEmail))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) RegisterByMobile(ctx context.Context, in *RegisterByMobileRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/api/user/registerByMobile"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserRegisterByMobile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *UserHTTPClientImpl) SendMobileVerifyCode(ctx context.Context, in *SendMobileVerifyCodeRequest, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/api/user/sendMobileVerifyCode"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationUserSendMobileVerifyCode))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
