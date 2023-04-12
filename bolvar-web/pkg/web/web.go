package web

import (
	api "bolvar-web/pkg/server"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

type bolvarServer struct {
	api.UnimplementedBolvarServiceServer
}

func (b bolvarServer) CreateEvent(ctx context.Context, request *api.CreateEventRequest) (*api.CreateEventReply, error) {
	//TODO implement me
	return &api.CreateEventReply{
		Message: fmt.Sprintf("get===>%s", request.Name),
	}, nil
}

func NewBolvarServer() *bolvarServer {
	return &bolvarServer{}
}

func Start(port int, webPort int) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {

	}

	s := grpc.NewServer()
	api.RegisterBolvarServiceServer(s, &bolvarServer{})
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%d", port),
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {

	}

	gwmux := runtime.NewServeMux()
	err = api.RegisterBolvarServiceHandler(context.Background(),
		gwmux,
		conn,
	)
	if err != nil {

	}
	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: gwmux,
	}
	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:%d\n", webPort)
	log.Fatalln(gwServer.ListenAndServe())
}
