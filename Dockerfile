# Build the port-exporter binary
FROM golang:1.9.2 as builder
ENV GOPATH /go
# Copy in the go src
WORKDIR /go/src/github.com/port-exporter
COPY pkg/    pkg/
COPY cmd/    cmd/
COPY vendor/ vendor/
COPY collector/ collector/
COPY Makefile Makefile

# Build
RUN make port-exporter

# Copy the controller-manager into a thin image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go/src/github.com/port-exporter/bin/port_exporter .
