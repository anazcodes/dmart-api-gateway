package client

import (
	"log"

	"github.com/anazibinurasheed/d-api-gateway/internal/auth-svc/pb"
	util "github.com/anazibinurasheed/d-api-gateway/internal/utils"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

func InitServiceClient(authSvcPort string) pb.AuthServiceClient {
	conn, err := grpc.Dial(authSvcPort, grpc.WithInsecure())
	if util.HasError(err) {
		log.Fatalln("failed to connect with auth service")
	}
	return pb.NewAuthServiceClient(conn)
}
