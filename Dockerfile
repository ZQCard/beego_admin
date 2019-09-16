FROM golang:latest

MAINTAINER Card "445864742@qq.com"

WORKDIR /usr/local/applications/www.fmg.ltd
ADD . /usr/local/applications/www.fmg.ltd

RUN go build -o fmg

EXPOSE 8080

ENTRYPOINT ["./fmg"]