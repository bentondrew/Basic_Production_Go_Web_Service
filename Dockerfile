FROM golang:1.11.0-alpine3.8 as builder
RUN adduser -D -g '' gouser
COPY source/ $GOPATH/src/web_service
WORKDIR $GOPATH/src/web_service
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags="-w -s" -o /go/bin/wiki

FROM scratch
LABEL maintainer="Benton Drew <benton.s.drew@drewantech.com>"
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --chown=gouser --from=builder /go/bin/ /web_service/
COPY --chown=gouser static/ /web_service/
USER gouser
WORKDIR /web_service/
EXPOSE 8080
ENTRYPOINT ["/web_service/service"]
