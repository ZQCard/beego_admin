FROM golang:latest

MAINTAINER Card "445864742@qq.com"

ENV GOPROXY=https://goproxy.io,direct

WORKDIR /usr/local/applications/taige.niu12.com

ADD . /usr/local/applications/taige.niu12.com

RUN go build -o taige

EXPOSE 8081

ENTRYPOINT ["./taige"]
