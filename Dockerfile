FROM golang:alpine

ENV GOPATH /go
ENV GO111MODULE on

RUN apk add --no-cache git
WORKDIR /flakbase
ADD . /flakbase
RUN go install

FROM alpine:latest

EXPOSE 9527
COPY --from=0 /go/bin/flakbase /usr/bin/flakbase

CMD flakbase serve -m /data/mongo.json
