package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	equities "github.com/dhananjayksharma/dkgosql-grpc-equities/equities"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// Implement the equities service (equities.equitiesServer interface)
type equitiesServer struct {
	equities.UnimplementedOrderServer
}

func (es *equitiesServer) ProcessOrder(stream equities.Order_ProcessOrderServer) error {
	for {
		// Get order for userid
		order, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("Received a order to process: %v", order)
		err = equities.ProcessOrder(&equities.OrderRequest{
			Userid:  order.Userid,
			Orderid: order.Orderid,
		})

		if err != nil {
			return err
		}
		log.Println("order: ", order)
	}
	return nil
}

func main() {
	// parse arguments from the command line
	// this lets us define the port for the server
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	// Check for errors
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// Instantiate the server
	s := grpc.NewServer()
	// Register server method (actions the server will do)
	equities.RegisterOrderServer(s, &equitiesServer{})
	// Register server method (actions the server will do)
	// TODO

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
