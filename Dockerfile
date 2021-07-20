FROM golang:1.16.6-alpine3.14

RUN mkdir /currencyapp

ADD . /currencyapp

WORKDIR /currencyapp

RUN go build -o koacurren .

CMD ["/currencyapp/koacurren"]

EXPOSE 10001