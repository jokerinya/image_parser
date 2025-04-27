package helpers

import (
	"fmt"
	"image"
	"image/jpeg"
	"math"
	"os"

	"golang.org/x/image/draw"
)

type Data struct {
	Filename string
	Img      image.Image
}

func LoadImage(filename string) (image.Image, error) {
	path := fmt.Sprintf("raw_images/%s", filename)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return img, nil
}

func ResizeImage(data *Data) error {
	outputPath := fmt.Sprintf("parsed_images/%s", data.Filename)
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	width := 300 // pixel, height will be automatic
	ratio := (float64)(data.Img.Bounds().Max.Y) / (float64)(data.Img.Bounds().Max.X)
	height := int(math.Round(float64(width) * ratio))

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(dst, dst.Rect, data.Img, data.Img.Bounds(), draw.Over, nil)

	err = jpeg.Encode(file, dst, nil)
	if err != nil {
		return err
	}

	return nil
}
