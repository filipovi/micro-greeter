FROM golang:1.10 AS builder
WORKDIR /go/src/github.com/filipovi/micro-greeter
COPY . /go/src/github.com/filipovi/micro-greeter
RUN go get github.com/micro/go-micro
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o greeter .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/filipovi/micro-greeter .
CMD ["./greeter"]

