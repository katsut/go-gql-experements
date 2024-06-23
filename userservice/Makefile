generate_grpc:
	protoc ./pb/service.proto --go_out=.  --go-grpc_out=require_unimplemented_servers=false:.

generate_dao:
	go run cmd/gormgen/gormgen.go
