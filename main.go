package main

import (
  "fmt"
  "errors"
  "net/http"
  "encoding/json"
)
var conversions = map[string]map[string]float64{
  "KSH":{
    "NGN":3,
    "GHS":3,
  },
  "NGN":{
    "KSH":12.1,
    "GHS":2.2,
  },
  "GHS":{
    "USD":4.23,
    "NGN":2.23,
  },
}
type ConvertRequest struct{
  From   string `json:"from"`
  To     string `json:"to"`
  Amount float64 `json:"amount"`
}

type Currency struct{
  Symbol   string `json:"symbol"`
  Amount   float64  `json:"amount"`
}
type ConvertResponse struct{
  Status int `json:"status"`
  Message string `json:"status_message"`
  Results Currency  `json:"results"`
}

func conversionHandler(w http.ResponseWriter,r *http.Request){
  switch r.Method {
    case "POST":
      var convertrequest ConvertRequest

      err := json.NewDecoder(r.Body).Decode(&convertrequest)
      if err != nil {
          http.Error(w, err.Error(), http.StatusBadRequest)
          return
      }
      w.Header().Set("Content-Type", "application/json; charset=utf-8")
      var resp ConvertResponse
      currency,err := convertrequest.Convert()
      if err != nil{
        resp.Status = -1
        resp.Message = err.Error()
        json.NewEncoder(w).Encode(resp)
        return;
      }
      resp.Status = 0
      resp.Message = "Query was successful"
      resp.Results = *currency
      json.NewEncoder(w).Encode(resp)
    default:
      fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}
func (req *ConvertRequest) Convert() (*Currency,error){
  rates, found := conversions[req.From]

  if !found{
    return nil, errors.New("No From")
  }
  rate, found := rates[req.To]
  if !found{
    return nil, errors.New("No To")
  }

  if req.Amount <= 0{
    return nil, errors.New("Bro No negative value")
  }

  var currency Currency
  currency.Symbol = req.To
  currency.Amount = req.Amount * rate
  return &currency,nil
}
func main(){
  mux := http.NewServeMux()
  fmt.Println("Listening on port :10001")
  mux.HandleFunc("/api/convert",conversionHandler)
  http.ListenAndServe(":10001",mux)
}
