package answerfromserv

import (
	"context"
	"gRPCcalculatorClient/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func GetAnswerFromServer(address string, par1 float32, par2 float32) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("not connection: %s", err)
	}
	defer conn.Close()
	c := api.NewDatabusServiceClient(conn)
	msg := api.SendRequest{
		Prm1: par1,
		Prm2: par2,
	}

	response, err := c.Send(context.Background(), &msg)
	if err != nil {
		log.Fatalf("error %s", err)
	}
	log.Println(response)
}
