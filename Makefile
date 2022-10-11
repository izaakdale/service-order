PROTO_DIR=pkg/model/order

run:
	SERVICE_ADDRESS=8080 SERVICE_DYNAMOTABLENAME=ordering-app SERVICE_DYNAMOREGION=eu-west-2 SERVICE_GRPCADDRESS=localhost:50001 go run main.go

gproto:
	protoc --proto_path=. --go_out=. --go_opt=paths=source_relative ${PROTO_DIR}/*.proto \
	 --go-grpc_out=. --go-grpc_opt=paths=source_relative ${PROTO_DIR}/*.proto