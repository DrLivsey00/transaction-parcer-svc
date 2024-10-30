FROM golang:1.20-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/DrLivsey00/transaction-parcer-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/transac-parser-svc /go/src/github.com/DrLivsey00/transaction-parcer-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/transac-parser-svc /usr/local/bin/transac-parser-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["transac-parser-svc"]
