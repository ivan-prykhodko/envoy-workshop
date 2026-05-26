package currency

import "fmt"

type ExchangeRateService struct {
}

func (s *ExchangeRateService) GetExchangeRate(currencyFrom, currencyTo string) (float64, error) {
	if currencyFrom != "UAH" {
		return 0, fmt.Errorf("unsupported currency: %s", currencyFrom)
	}
	return s.findExchangeRate(currencyFrom, currencyTo)
}

func (s *ExchangeRateService) findExchangeRate(currencyFrom, currencyTo string) (float64, error) {
	rateFrom, err := s.findRate(currencyFrom)
	if err != nil {
		return 0, fmt.Errorf("failed to find rate for currency %s: %v", currencyFrom, err)
	}
	rateTo, err := s.findRate(currencyTo)
	if err != nil {
		return 0, fmt.Errorf("failed to find rate for currency %s: %v", currencyTo, err)
	}
	// TODO: check if rateTo != 0
	return rateFrom / rateTo, nil
}

func (s *ExchangeRateService) findRate(currency string) (float64, error) {
	var rates = s.allRates()
	rate, ok := rates[currency]
	if !ok {
		return 0, fmt.Errorf("unsupported currency: %s", currency)
	}
	return rate, nil
}

func (s *ExchangeRateService) allRates() map[string]float64 {
	var rates = make(map[string]float64)
	rates["UAH"] = 1
	rates["USD"] = 44.01
	rates["EUR"] = 52.08
	return rates
}
