package exchange

import (
	"errors"
	"fmt"
	"time"
)

type ConvertRequest struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}

type Currency struct {
	Symbol    string  `json:"symbol"`
	Amount    float64 `json:"amount"`
	TimeStamp int64   `json:"timestamp"`
}
type ConvertResponse struct {
	Status  int      `json:"status"`
	Message string   `json:"status_message"`
	Results Currency `json:"results"`
}

func (req *ConvertRequest) Convert() (*Currency, error) {
	rates, found := Conversions[req.From]
	if !found {
		return nil, fmt.Errorf("conversions from currency %s are not supported at the time", req.From)
	}

	rate, found := rates[req.To]
	if !found {
		return nil, fmt.Errorf("conversions of currency %s to %s is not supported at the time", req.From, req.To)
	}

	if req.Amount <= 0 {
		return nil, errors.New("conversion rates only occur whit positive values greater than zero")
	}

	var currency Currency
	currency.Symbol = req.To
	currency.Amount = req.Amount * rate
	currency.TimeStamp = time.Now().Unix()
	return &currency, nil
}

func (c *Currency) GetAmount() float64 {
	return c.Amount
}
