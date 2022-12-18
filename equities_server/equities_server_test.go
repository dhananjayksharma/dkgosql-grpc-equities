package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net"
	"testing"

	gRPGEquities "github.com/dhananjayksharma/dkgosql-grpc-equities/equities"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/test/bufconn"
)

func server(ctx context.Context) (gRPGEquities.OrderClient, func()) {
	buffer := 101024 * 1024
	lis := bufconn.Listen(buffer)

	baseServer := grpc.NewServer()
	gRPGEquities.RegisterOrderServer(baseServer, &equitiesServer{})
	go func() {
		if err := baseServer.Serve(lis); err != nil {
			log.Printf("error serving server: %v", err)
		}
	}()

	conn, err := grpc.DialContext(ctx, "",
		grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return lis.Dial()
		}), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("error connecting to server: %v", err)
	}

	closer := func() {
		err := lis.Close()
		if err != nil {
			log.Printf("error closing listener: %v", err)
		}
		baseServer.Stop()
	}

	client := gRPGEquities.NewOrderClient(conn)

	return client, closer
}

// TestTelephoneServer_SendMessage
func TestTelephoneServer_ProcessOrder(t *testing.T) {
	ctx := context.Background()
	client, closer := server(ctx)
	defer closer()

	type expectation struct {
		out []*gRPGEquities.OrderResponse
		err error
	}

	tests := map[string]struct {
		in       []*gRPGEquities.OrderRequest
		expected expectation
	}{
		"Must_Success": {
			in: []*gRPGEquities.OrderRequest{
				{
					Orderid: "13",
					Userid:  "12",
				},
				{
					Orderid: "131",
					Userid:  "122",
				},
				{
					Orderid: "123456",
					Userid:  "122",
				},
			},
			expected: expectation{
				out: []*gRPGEquities.OrderResponse{
					{
						Orderid: "13",
						Userid:  "12",
					},
					{
						Orderid: "131",
						Userid:  "122",
					},
					{
						Orderid: "123456",
						Userid:  "122",
					},
				},
				err: nil,
			},
		},
	}

	for scenario, tt := range tests {
		t.Run(scenario, func(t *testing.T) {
			outClient, err := client.ProcessOrder(ctx)
			if err != nil {
				t.Errorf("Error for client:%v", err)
			}

			for _, v := range tt.in {
				if err := outClient.Send(v); err != nil {
					t.Errorf("Err -> %q", err)
				}
			}

			if err := outClient.CloseSend(); err != nil {
				t.Errorf("Err -> %q", err)
			}

			var outs []*gRPGEquities.OrderResponse
			for {
				o, err := outClient.Recv()
				if errors.Is(err, io.EOF) {
					break
				}
				outs = append(outs, o)
			}

			for i, o := range outs {
				if !assert.Equal(t, o.GetOrderid(), tt.expected.out[i].GetOrderid()) {
					t.Errorf("Out -> \nWant: %q\nGot : %q", tt.expected.out[i].GetOrderid(), o.GetOrderid())
				}
			}

		})
	}
}
