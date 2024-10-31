FROM golang:1.23-alpine as buildbase

RUN apk add git build-base


WORKDIR /go/src/github.com/DrLivsey00/transaction-parcer-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/transac-parser-svc /go/src/github.com/DrLivsey00/transaction-parcer-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/transac-parser-svc /usr/local/bin/transac-parser-svc
COPY config.local.yaml /usr/local/bin/config/config.local.yaml
COPY nginx.conf /etc/nginx/nginx.conf
COPY entry.sh /usr/local/bin/entrypoint.sh

RUN chmod +x /usr/local/bin/entrypoint.sh
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["/usr/local/bin/entrypoint.sh"]
