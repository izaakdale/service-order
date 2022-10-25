FROM alpine
WORKDIR /build
COPY service-order .

# these should be in compose/k8s
# ENV SERVICE_ADDRESS=80
# ENV SERVICE_DYNAMOTABLENAME=ordering-app
# ENV SERVICE_DYNAMOREGION=eu-west-2
# ENV SERVICE_GRPCADDRESS=localhost:50001

CMD ["./service-order"]