FROM golang:1.20-alpine as builder
WORKDIR /

COPY . ./
RUN go mod download


RUN go build -o /service-order

FROM alpine
COPY --from=builder /service-order .

EXPOSE 80
CMD [ "/service-order" ]