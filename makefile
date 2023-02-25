run:
	QUEUE_URL=http://localhost:4566/000000000000/order-placed-queue \
	AWS_REGION=eu-west-2 \
	AWS_ENDPOINT=http://localhost:4566 \
	TOPIC_ARN=arn:aws:sns:eu-west-2:000000000000:order-stored-events \
	TABLE_NAME=orders \
	GRPC_HOST=localhost \
	GRPC_PORT=50002 \
	go run .

PROTO_DIR=pkg/proto/order
gproto:
	protoc \
	--proto_path=. \
	--go_out=. \
	--go_opt=paths=source_relative \
	${PROTO_DIR}/*.proto \
	 --go-grpc_out=. \
	 --go-grpc_opt=paths=source_relative \
	 ${PROTO_DIR}/*.proto