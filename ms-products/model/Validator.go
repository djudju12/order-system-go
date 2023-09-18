package model

import (
	"log"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var ValidPrice validator.Func = func(fl validator.FieldLevel) bool {
	if price, ok := fl.Field().Interface().(string); ok {
		return isValidPrice(price)
	}

	return false
}

func isValidPrice(price string) bool {
	// regex that checks if the price is valid (decimal(12, 2))
	pattern := `^\d{1,10}(\.\d{1,2})?$`
	matched, err := regexp.MatchString(pattern, price)
	if err != nil {
		log.Fatal("error on compile regex:", err)
	}

	return matched
}

var ValidStatus validator.Func = func(fl validator.FieldLevel) bool {
	if status, ok := fl.Field().Interface().(string); ok {
		return isValidStatus(status)
	}

	return false
}

func isValidStatus(status string) bool {
	switch status {
	case ProductStatusAvailable, ProductStatusOutOfStock:
		return true
	}

	return false
}
