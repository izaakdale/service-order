run:
	QUEUE_URL=http://localhost:4566/000000000000/order-placed-queue \
	AWS_REGION=eu-west-2 \
	AWS_ENDPOINT=http://localhost:4566 \
	TOPIC_ARN=arn:aws:sns:eu-west-2:000000000000:order-stored-events \
	go run .