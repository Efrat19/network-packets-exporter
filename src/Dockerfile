FROM golang:1.13

WORKDIR /go/src/app
COPY .. .

RUN apt-get update
RUN apt-get install -y libpcap-dev

RUN go get -d -v ./...
RUN go build .

# Interface to read packets from
ENV IFACE=eth0

# Exposed metrics port
ENV METRICS_PORT=9717

CMD ["./app"]

EXPOSE $METRICS_PORT