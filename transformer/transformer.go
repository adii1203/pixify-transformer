package transformer

import (
	"strconv"
	"strings"

	"github.com/adii1203/pixify-transformer/utils"
	"github.com/h2non/bimg"
)

type TransformFunc func([]byte, string) ([]byte, error)

var transformations = map[string]TransformFunc{
	"wh": resizeWidthHeight,
}

func ApplyTransformations(buf []byte, params map[string]string) ([]byte, error) {
	var err error
	img := bimg.NewImage(buf)
	metaData, _ := img.Metadata()

	width, hasWidth := utils.ParseDimension(params["w"], metaData.Size.Width)
	height, hasHeight := utils.ParseDimension(params["h"], metaData.Size.Height)

	if hasWidth && !hasHeight {
		height = int(float64(metaData.Size.Height) * (float64(width) / float64(metaData.Size.Width)))
		params["wh"] = strconv.Itoa(width) + "-" + strconv.Itoa(height)
	} else if hasHeight && !hasWidth {
		width = int(float64(metaData.Size.Width) * (float64(height) / float64(metaData.Size.Height)))
		params["wh"] = strconv.Itoa(width) + "-" + strconv.Itoa(height)
	} else {
		params["wh"] = strconv.Itoa(width) + "-" + strconv.Itoa(height)
	}

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

func resizeWidthHeight(buf []byte, value string) ([]byte, error) {

	dimensions := strings.Split(value, "-")

	height, err := strconv.ParseFloat(dimensions[1], 64)
	if err != nil {
		return buf, nil
	}
	width, err := strconv.ParseFloat(dimensions[0], 64)
	if err != nil {
		return buf, nil
	}

	return bimg.NewImage(buf).ResizeAndCrop(int(width), int(height))
}
