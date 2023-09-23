package authsvc

import (
	"log"

	"github.com/anazibinurasheed/dmart-api-gateway/pkg/auth-svc/pb"
	util "github.com/anazibinurasheed/dmart-api-gateway/pkg/util"
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
