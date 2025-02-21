package utils

import (
	"strconv"
	"strings"
)

func ExtractTransformationsOptions(path string) map[string]string {
	// path ==> tr:w=10,h=40
	optString := strings.Split(path, ":")[1]
	optStringArray := strings.Split(optString, ",")
	m := make(map[string]string)
	for _, opt := range optStringArray {
		kv := strings.Split(opt, "=")
		m[kv[0]] = kv[1]
	}
	return m
}

func ParseDimension(value string, originalSize int) (int, bool) {
	if value == "" {
		return 0, false
	}

	dim, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return 0, false
	}

	if dim > 0 && dim <= 1 {
		valueInPercent := float64(dim) * 100
		dim = (valueInPercent * float64(originalSize)) / 100
	}

	return int(dim), true
}
