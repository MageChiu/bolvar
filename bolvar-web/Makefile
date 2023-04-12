gen-grpc:
	protoc -I ./proto \
    --go_out ./pkg/server --go_opt=paths=source_relative \
    --go-grpc_out ./pkg/server  --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./pkg/server --grpc-gateway_opt paths=source_relative \
    --openapiv2_out ./openapi/openapiv2 --openapiv2_opt logtostderr=true \
    --swagger_out ./openapi/swagger --swagger_opt logtostderr=true \
    --gofast_out ./pkg/fast \
    ./proto/bolvar_web_server.proto