FROM ghcr.io/thanhpp/scmsrv-base:latest AS builder

# set ENV variables
ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/github.com/thanhpp/scm-tool

RUN apk add tzdata
RUN cp -r /usr/share/zoneinfo /etc/localtime

COPY pkg pkg/
COPY cmd cmd/
COPY internal internal/
COPY go.mod go.mod

RUN go mod tidy

WORKDIR /go/src/github.com/thanhpp/scm-tool/cmd/scmsrv

RUN go build -v -o scmsrv


# ----------------------------------------------------------------------------

FROM alpine

COPY --from=builder ["/go/src/github.com/thanhpp/scm-tool/cmd/scmsrv/scmsrv", "scmsrv"]
COPY --from=builder ["/go/src/github.com/thanhpp/scm-tool/cmd/scmsrv/config.yml", "config.yml"]

EXPOSE 10000

CMD ["./scmsrv"]