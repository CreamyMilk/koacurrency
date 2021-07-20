FROM golang

RUN mkdir /currencyapp

ADD . /currencyapp

WORKDIR /currencyapp

RUN go build -o koacurren .

CMD ["/currencyapp/koacurren"]

EXPOSE 10001