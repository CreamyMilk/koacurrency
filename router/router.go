package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/CreamyMilk/koacurrency/exchange"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/convert", conversionHandler)
	mux.HandleFunc("/api/rates", getConversionTable)
}

func conversionHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var convertrequest exchange.ConvertRequest

		err := json.NewDecoder(r.Body).Decode(&convertrequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		var resp exchange.ConvertResponse
		currency, err := convertrequest.Convert()
		if err != nil {
			resp.Status = -1
			resp.Message = err.Error()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(resp)
			return
		}

		resp.Status = 0
		resp.Message = "Conversion was successful"
		resp.Results = *currency
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(resp)
	default:
		w.WriteHeader(http.StatusFound)
		fmt.Fprintf(w, "Sorry, POST method is only supported.")
	}
}

func getConversionTable(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(exchange.Conversions)
}
