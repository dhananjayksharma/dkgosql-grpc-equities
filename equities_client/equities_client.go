package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"

	"github.com/dhananjayksharma/dkgosql-grpc-equities/equities"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := equities.NewOrderClient(conn)

	stream, err := client.ProcessOrder(context.Background())

	if err != nil {
		log.Fatalln("Opening stream", err)
	}
	for i := 0; i < 10; i++ {
		orderid := uuid.New()
		userid := uuid.New()
		if err := stream.Send(&equities.OrderRequest{
			Orderid: orderid.String(),
			Userid:  userid.String(),
		}); err != nil {
			log.Fatalln("Send", err)
		}
		log.Printf("Sending userid, orderid:%v, %v", userid, orderid)
	}
	if err := stream.CloseSend(); err != nil {
		log.Fatalln("CloseSend", err)
	}

	for {
		readRow, err := stream.Recv()
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalln("Recv", err)
		}
		fmt.Printf("Processed GetUserid:%s, GetOrderid:%s, Status:%t, Time: %v\n", readRow.GetUserid(), readRow.GetOrderid(), readRow.GetStatus(), readRow.GetNewupdateddt())
	}

}
