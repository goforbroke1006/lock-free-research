FROM golang:1.15 AS builder

WORKDIR /code

COPY ./go.mod ./
RUN go mod download

COPY ./ ./
RUN ls -lah
RUN CGO_ENABLED=0 go build

#
#

FROM debian:stretch

RUN apt-get update
RUN apt-get install -y curl

WORKDIR /app

COPY --from=builder /code/lock-free-research ./

EXPOSE 8080
EXPOSE 10000

HEALTHCHECK --timeout=10s --interval=5s --retries=10 CMD curl --fail http://127.0.0.1:8080/metrics || exit 1

ENTRYPOINT [ "./lock-free-research" ]
