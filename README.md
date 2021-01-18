# lock-free-research

### Benchmarks

```bash
go build
./lock-free-research util counter-spam --addr=192.168.0.211:10001 --concurrent=10000 --interval=50ms
./lock-free-research util counter-spam --addr=192.168.0.211:10002 --concurrent=10000 --interval=50ms
# where 192.168.0.211 is performance stand IP

```

### Report

#### Performance stand (bare-metal)

* AMD Ryzen 9 3900 12-Core Processor
* RAM 32 Gb
* Limits - defined in **docker-compose.yaml**

#### Metrics

<img src="./report.png" alt="Grafana metrics sample"/>

### Useful links

* [LB for compose](https://pspdfkit.com/blog/2018/how-to-use-docker-compose-to-run-multiple-instances-of-a-service-in-development/)
* [Nginx LB for TCP & UPD](https://docs.nginx.com/nginx/admin-guide/load-balancer/tcp-udp-load-balancer/)
