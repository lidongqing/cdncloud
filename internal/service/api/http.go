package api

import "net/http"

type HttpService struct {
	request *http.Request
}

func NewHttpService(request *http.Request) *HttpService {
	return &HttpService{
		request: request,
	}
}
