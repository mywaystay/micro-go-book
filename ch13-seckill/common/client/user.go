package client

import (
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/pb"
	endpts "github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/user-service/endpoint"
	"github.com/keets2012/Micro-Go-Pracrise/ch13-seckill/user-service/service"
	"google.golang.org/grpc"
)

func UserCheck(conn *grpc.ClientConn) service.Service {
	var ep = grpctransport.NewClient(conn,
		"pb.UserService",
		"Check",
		EncodeGRPCUserRequest,
		DecodeGRPCUserResponse,
		pb.UserResponse{},
	).Endpoint()

	userEp := endpts.UserEndpoints{
		UserEndpoint: ep,
	}
	return userEp
}