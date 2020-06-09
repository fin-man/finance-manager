#first stage - builder
FROM golang:1.13.0-stretch as builder

COPY . /finance-manager
WORKDIR /finance-manager/server

ENV GO111MODULE=on

RUN CGO_ENABLED=0 GOOS=linux go build -o FinMan 


#second stage 
FROM alpine:latest

RUN apk add --no-cache tzdata

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /root/

COPY --from=builder /finance-manager/server .

EXPOSE 8080
EXPOSE 6379

CMD ["./FinMan"]