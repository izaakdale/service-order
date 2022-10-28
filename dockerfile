FROM alpine
WORKDIR /build
COPY service-order .

CMD ["./service-order"]