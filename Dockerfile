FROM golang:1.13

WORKDIR /go/src/app
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

# Interface to read packets from
ENV IFACE=eth0

# Exposed metrics port
ENV METRICS_PORT=9717

CMD ["app"]

EXPOSE $METRICS_PORT