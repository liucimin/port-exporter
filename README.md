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



## Building and running

### Prerequisites:

- [Go compiler](https://golang.org/dl/) or Docker

- RHEL/CentOS: `glibc-static` package.

  

### Building with Go Env:

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

### Building with Docker:

```
go get github.com/liucimin/port-exporter
cd ${GOPATH-$HOME/go}/src/github.com/liucimin/port-exporter
make docker-build
```

### Running  with container:

```
docker run -it  --pid="host" --network="host"   --security-opt seccomp=unconfined  --privileged -v /var/run/:/var/run/ port-exporter:1.0  sh

 ./port_exporter  --v=5  --logtostderr=true
```

Then we can get the metrics from the port_exporter:

```
curl container:9103/metrics
```



## Running tests

```
make test
```


## Kubernetes Support
