package api

import "github.com/go-playground/validator/v10"

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return isCurrencySupported(currency)
	}
	return false

}

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
