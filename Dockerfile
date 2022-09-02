FROM golang:1.18 AS builder

WORKDIR /go/src/github.com/aboodtbr/server
COPY . .

RUN go mod download
RUN go build -o /go/bin/server

FROM bubuntux/nordvpn:v3.10.0-1-1

HEALTHCHECK --interval=5m --timeout=20s --start-period=1m \
  CMD if test $( curl -m 10 -s https://api.nordvpn.com/vpn/check/full | jq -r '.["status"]' ) = "Protected" ; then exit 0; else exit 1; fi

COPY --from=builder /go/bin/server .
RUN chmod 777 server

COPY start.sh .
RUN chmod 777 start.sh

EXPOSE 1080

ENTRYPOINT ["./start.sh"]