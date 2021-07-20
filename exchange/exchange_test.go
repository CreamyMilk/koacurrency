package exchange

import (
	"testing"
)

func TestConvert(t *testing.T) {

	tt := []struct {
		testName   string
		from       string
		to         string
		amount     float64
		expected   float64
		isPossible bool
	}{
		{"KSH to NGN standard", "KSH", "NGN", 100, 300, true},
		{"NGN to GSH standard", "NGN", "GHS", 200, 2, true},
		{"KSH to NGN decimals", "KSH", "NGN", 0.01, 0.03, true},
		{"NGN to GHS decimals", "NGN", "GHS", 15.50, 0.155, true},
		{"Negative Amount", "KSH", "GHS", 100, 0, false},
		{"Zero Amount", "NGN", "KSH", 100, 0, false},
		{"KSH to NONE exstant currency", "KSH", "UNKNOWN", 100, 100, false},
		{"NONE existant to existant currency", "UNKNOWN", "KSH", 100, 100, false},
	}

	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			req := ConvertRequest{From: tc.from, To: tc.to, Amount: tc.amount}
			result, err := req.Convert()
			if err != nil && tc.isPossible {
				t.Errorf("Cannot perform conversion because %v", err)
			}
			if tc.isPossible {
				finalRate := result.GetAmount()
				if finalRate != tc.expected {
					t.Errorf("Expected Currency Amount to be %v but got %v", tc.expected, finalRate)
				}
			}
		})
	}
}
