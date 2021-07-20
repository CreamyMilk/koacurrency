package exchange

//Conversions holds the to and from rates of each currency
//It is accessed as follows
//
//conversion_rate := Conversions[from][to]
var Conversions = map[string]map[string]float64{
	"KSH": {
		"NGN": 3,
		"GHS": 3,
	},
	"NGN": {
		"GHS": 0.01,
		"KSH": 0.2,
	},
	"GHS": {
		"KSH": 18.2,
		"NGN": 69.3,
	},
}
