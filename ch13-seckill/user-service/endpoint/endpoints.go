package endpoint

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	cli "github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/common/client/user_client"
	"github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/user-service/service"
)

// UserEndpoint define endpoint
func (u *cli.UserEndpoints) Check(ctx context.Context, username string, password string) (bool, error) {
	//ctx := context.Background()
	resp, err := u.UserEndpoint(ctx, UserRequest{
		Username: username,
		Password: password,
	})
	response := resp.(UserResponse)
	err = errors.New("bad request!")
	return response.Result, err
}

func (ue cli.UserEndpoints) HealthCheck() bool {
	return false
}

var (
	ErrInvalidRequestType = errors.New("invalid username, password")
)

// UserRequest define request struct
type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// UserResponse define response struct
type UserResponse struct {
	Result bool  `json:"result"`
	Error  error `json:"error"`
}

//  make endpoint
func MakeUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)

		var (
			username, password string
			res                bool
			calError           error
		)

		username = req.Username
		password = req.Password

		res, calError = svc.Check(ctx, username, password)
		if calError != nil {
			return UserResponse{Result: false, Error: calError}, nil
		}
		return UserResponse{Result: res, Error: calError}, nil
	}
}

// HealthRequest 健康检查请求结构
type HealthRequest struct{}

// HealthResponse 健康检查响应结构
type HealthResponse struct {
	Status bool `json:"status"`
}

// MakeHealthCheckEndpoint 创建健康检查Endpoint
func MakeHealthCheckEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		status := svc.HealthCheck()
		return HealthResponse{status}, nil
	}
}
