package web

import (
	api "bolvar-web/pkg/server"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/grpclog"
	"net/http"
	"strings"
)

func StartGin(webPort int, grpcPort int) {
	//grpclog.SetLoggerV2(grpclog.NewLoggerV2())
	grpclog.Infof("start gin")
	var r = gin.Default()
	grpcServer := grpc.NewServer()
	api.RegisterBolvarServiceServer(grpcServer, &bolvarServer{})

	r.Use(func(context *gin.Context) {
		if context.Request.ProtoMajor == 2 &&
			strings.HasPrefix(context.GetHeader("Content-Type"), "application/grpc") {
			context.Status(http.StatusOK)
			grpcServer.ServeHTTP(context.Writer, context.Request)
			//
			context.Abort()
			return
		}
		context.Next()
	})
	gwmux := runtime.NewServeMux()
	ctx := context.Background()
	endpoint := fmt.Sprintf("127.0.0.1:%d", webPort)

	dopts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	if err := api.RegisterBolvarServiceHandlerFromEndpoint(ctx, gwmux, endpoint, dopts); err != nil {
		grpclog.Fatalf("Failed to register gw server: %v\n", err)
	}
	r.Any("/v1/event/create", func(c *gin.Context) {
		gwmux.ServeHTTP(c.Writer, c.Request)
	})
	h2Handle := h2c.NewHandler(r.Handler(), &http2.Server{})
	server := &http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%d", webPort),
		Handler: h2Handle,
	}
	// 启动http服务
	server.ListenAndServe()
}
