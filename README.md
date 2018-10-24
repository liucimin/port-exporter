# port-exporter

Prometheus exporter for port metrics.written in Go with pluggable metric collectors.

## Collectors

There is varying support for collectors on each operating system. The tables
below list all existing collectors and the supported systems.

### Enabled by default

| Name                | Description                                          | OS    |
| ------------------- | ---------------------------------------------------- | ----- |
| ovs_db              | Exposes ovs ports statistics from ovsdb.             | Linux |
| container_interface | Exposes container interfaces state from `container`. | Linux |

### Disabled by default

| Name                | Description                                          | OS    |
| ------------------- | ---------------------------------------------------- | ----- |
| ovs_db              | Exposes ovs ports statistics from ovsdb.             | Linux |
| container_interface | Exposes container interfaces state from `container`. | Linux |



## Building and running

Prerequisites:

- [Go compiler](https://golang.org/dl/) or Docker

- RHEL/CentOS: `glibc-static` package.

  

Building with Go Env:

```
go get github.com/liucimin/port-exporter
cd ${GOPATH-$HOME/go}/src/github.com/liucimin/port-exporter
make
./port_exporter <flags>
```

To see all available configuration flags:

```
./port_exporter -h
```

Building with Docker:

```
go get github.com/liucimin/port-exporter
cd ${GOPATH-$HOME/go}/src/github.com/liucimin/port-exporter
make with-docker
```

## Running tests

```
make test
```


