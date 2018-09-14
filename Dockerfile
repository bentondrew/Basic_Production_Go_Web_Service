FROM golang:1.11.0-alpine3.8 as builder
RUN adduser -D -g '' gouser
COPY source/ $GOPATH/src/service/
WORKDIR $GOPATH/src/service/
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/main

FROM scratch
LABEL maintainer="Benton Drew <benton.s.drew@drewantech.com>"
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --chown=gouser --from=builder /go/bin/ /service/
COPY tls/ /tls/
ENV CERT_FILE /tls/local.localhost.cert
ENV KEY_FILE /tls/local.localhost.key
ENV SERVICE_ADDR ":8080"
USER gouser
WORKDIR /service/
EXPOSE 8080
ENTRYPOINT ["/service/main"]
