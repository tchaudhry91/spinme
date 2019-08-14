package spin

import (
	"errors"
	"strings"
)

func SpinnerFrom(service string) (SpinnerFunc, error) {
	switch s := strings.ToLower(service); s {
	case "mongo":
		return Mongo, nil
	default:
		return Generic, errors.New("Failed to find given service")
	}
}
