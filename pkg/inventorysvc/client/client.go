package client

import (
	"log"

	"github.com/anazibinurasheed/d-api-gateway/internal/inventorysvc/pb"
	"github.com/anazibinurasheed/d-api-gateway/internal/util"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.InventoryServiceClient
}

func InitServiceClient(inventorySvcPort string) pb.InventoryServiceClient {
	conn, err := grpc.Dial(inventorySvcPort, grpc.WithInsecure())
	if util.HasError(err) {
		log.Fatalln("failed to connect with inventory service")
	}
	return pb.NewInventoryServiceClient(conn)
}
