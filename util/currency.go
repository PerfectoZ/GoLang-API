package util

const (
	USD = "USD"
	INR = "INR"
	CAD = "CAD"
	EUR = "EUR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, INR, CAD, EUR:
		return true
	}
	return false
}
