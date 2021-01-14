FROM golang:1.15 AS builder

WORKDIR /code

COPY ./go.mod ./
RUN go mod download

COPY ./ ./
RUN ls -lah
RUN CGO_ENABLED=0 go build

#
#

FROM scratch

WORKDIR /app

COPY --from=builder /code/lock-free-research ./

ENTRYPOINT [ "./lock-free-research" ]
