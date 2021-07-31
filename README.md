# koacurrency
Minimal Currency exchange

To run the service you can run

```bash
$ PORT=:10001 go run main.go
```

To run unit tests run .
```bash
$ go test -v ./exchange
```

_To issue a conversion request_
```bash
curl --header "Content-Type: application/json"   --request POST   --data '
{
    "from":"KSH",
    "to":"NGN",
    "amount": 1000.5
}' http://localhost:10001/api/convert
```



_Get current conversion rates_

```bash
curl http://localhost:10001/api/rates
```

_Get current conversion rates of a particular currency_

```bash
curl localhost:10001/api/rates?currency="KSH"
```
----
> _You can also run the service using docker_
> ```bash
> $ docker build -t koacurrency .
> $ docker run -p 10001:10001 -it koacurrency 
> ```


