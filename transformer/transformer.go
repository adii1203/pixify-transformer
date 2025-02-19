package transformer

import (
	"strconv"

	"github.com/h2non/bimg"
)

type TransformFunc func([]byte, string) ([]byte, error)

var transformations = map[string]TransformFunc{
	"w": resizeWidth,
	"h": resizeHeight,
}

func ApplyTransformations(buf []byte, params map[string]string) ([]byte, error) {
	var err error
	for key, value := range params {
		if trf, ok := transformations[key]; ok {
			buf, err = trf(buf, value)
			if err != nil {
				return nil, err
			}
		}
	}
	return buf, nil
}

func resizeWidth(buf []byte, value string) ([]byte, error) {
	width, err := strconv.Atoi(value)
	if err != nil {
		return buf, nil
	}
	return bimg.NewImage(buf).Resize(width, 0)
}

func resizeHeight(buf []byte, value string) ([]byte, error) {
	height, err := strconv.Atoi(value)
	if err != nil {
		return buf, nil
	}
	return bimg.NewImage(buf).Resize(0, height)
}
