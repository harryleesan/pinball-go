FROM golang:1.9.2-stretch as builder
RUN apt-get update \
	&& apt-get -y install ca-certificates
WORKDIR /go/src/app
COPY . .
RUN go test
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
RUN chmod +x ./main

FROM scratch
MAINTAINER Harry Lee <harryl@vatit.com>
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/app/main /opt/app/main
WORKDIR /opt/app
CMD ["./main"]

