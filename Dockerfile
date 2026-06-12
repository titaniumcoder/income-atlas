ARG GO_VERSION=1.26-bullseye
FROM golang:${GO_VERSION}-bookworm as builder

WORKDIR /usr/src/app
COPY go.mod go.sum Makefile ./
RUN make install-generate
RUN go mod download && go mod verify

COPY . .

RUN make build-docker

FROM debian:bookworm

COPY --from=builder /income-atlas /usr/local/bin/
CMD ["income-atlas"]
