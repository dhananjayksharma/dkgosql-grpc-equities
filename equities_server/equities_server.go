package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	gRPGEquities "github.com/dhananjayksharma/dkgosql-grpc-equities/equities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

// Implement the equities service (equities.equitiesServer interface)
type equitiesServer struct {
	gRPGEquities.UnimplementedOrderServer
}

func (es *equitiesServer) ProcessOrder(stream gRPGEquities.Order_ProcessOrderServer) error {
	// ctx := stream.Context()
	for {
		// select {
		// case <-ctx.Done():
		// 	return ctx.Err()
		// default:
		// }
		var validationError *status.Status
		// Reading stream Request
		order, err := stream.Recv()
		if err == io.EOF {
			log.Printf("End of receive @time:%v", time.Now())
			break
		}
		if err != nil {
			log.Printf("ProcessOrder receive error %v", err)
			return err
		}

		// log.Printf("Received a order to process: %v", order)
		request := &gRPGEquities.OrderRequest{
			Userid:     order.Userid,
			Orderid:    order.Orderid,
			Quantity:   order.Quantity,
			Allowedqty: order.Allowedqty,
			Ordertype:  order.Ordertype,
		}
		// var err error
		if order.Quantity > order.Allowedqty {
			validationError = status.Newf(
				codes.OutOfRange,
				"order quantity exceeds max allowed quantity",
			)
			validationError, err = validationError.WithDetails(order)
			if err != nil {
				fmt.Errorf("unable to process this order to error", "error", err)
			}
		}

		// if a validationError return error
		if validationError != nil {
			stream.Send(&gRPGEquities.StreamingOrderResponse{Message: &gRPGEquities.StreamingOrderResponse_Error{
				Error: validationError.Proto(),
			}})
			continue
		}

		pQty, status, err := gRPGEquities.ProcessOrder(request)
		response := &gRPGEquities.StreamingOrderResponse{
			Message: &gRPGEquities.StreamingOrderResponse_Orderresponse{
				&gRPGEquities.OrderResponse{
					Orderid:                request.Orderid,
					Userid:                 request.Userid,
					Quantity:               request.Quantity,
					ProcessedQuantity:      pQty,
					NotProcessedQuantity:   request.Quantity - pQty,
					Status:                 status,
					Newupdateddt:           time.Now().String(),
					Orderprocessedupdatedt: timestamppb.Now(),
				},
			},
		}
		fmt.Println("Sending Response:", response)

		// Sending stream Reponse
		if err := stream.Send(response); err != nil {
			log.Printf("send error %v", err)
		}

		if err != nil {
			log.Printf("SERVER: error calling equities.ProcessOrder:  error %v", err)
			return err
		}

	}
	log.Println("Server equities.ProcessOrder quit")
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
	gRPGEquities.RegisterOrderServer(s, &equitiesServer{})
	// Register server method (actions the server will do)
	// TODO
	reflection.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
