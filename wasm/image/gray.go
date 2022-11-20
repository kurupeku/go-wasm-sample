package image

import (
	"bytes"
	"encoding/base64"
	"image"
	"image/png"
)

func GrayScale(imgStr string) (string, error) {
	b, err := base64.StdEncoding.DecodeString(imgStr)
	if err != nil {
		return "", err
	}

	img, err := png.Decode(bytes.NewReader(b))
	if err != nil {
		return "", err
	}

	bounds := img.Bounds()
	dest := image.NewGray(bounds)
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			dest.Set(x, y, img.At(x, y))
		}
	}

	rb := bytes.NewBuffer([]byte{})
	png.Encode(rb, dest)

	return base64.StdEncoding.EncodeToString(rb.Bytes()), nil
}
