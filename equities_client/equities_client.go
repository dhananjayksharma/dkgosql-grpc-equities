package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"time"

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
	// Set up a connection to the server.
	// TODO
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := equities.NewOrderClient(conn)

	// Define the context
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.ProcessOrder(ctx)

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
		_, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln("Recv", err)
		}
		fmt.Println("Processed GetUserid, GetOrderid")
	}
}
