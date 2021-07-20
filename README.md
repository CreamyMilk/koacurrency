# koacurrency
Minimal Currency exchange

To run the service you can run

```bash
$ PORT=:10001 go run main.go
```

To issue a conversion request
```bash
curl --header "Content-Type: application/json"   --request POST   --data '
{
    "from":"KSH",
    "to":"NGN",
    "amount": 12222.5
}' http://localhost:10001/api/convert
```

Get current conversion rates

```bash
curl http://localhost:10001/api/rates
```


