package utils

const (
	USD = "USD"
	INR = "INR"
	EUR = "EUR"
)

func isCurrencySupported(currency string) bool {
	switch currency {
	case USD, INR, EUR:
		return true
	}
	return false
}
