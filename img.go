package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"os"

	"golang.org/x/image/draw"
)

type Img struct {
	Filename string
	Data     image.Image
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

func (img *Img) Resize() {
	// https://roeber.dev/posts/resize-an-image-in-go/
	width := 300 // pixel, height will be automatic
	ratio := (float64)(img.Data.Bounds().Max.Y) / (float64)(img.Data.Bounds().Max.X)
	height := int(math.Round(float64(width) * ratio))

	dst := image.NewRGBA(image.Rect(0, 0, width, height))
	draw.NearestNeighbor.Scale(dst, dst.Rect, img.Data, img.Data.Bounds(), draw.Over, nil)

	img.Data = dst
}

func (img *Img) GrayScale() {
	// https://www.imager200.io/blog/grayscaling-image-golang/
	target := image.NewRGBA64(img.Data.Bounds())
	i := 0
	for i < img.Data.Bounds().Max.Y {
		j := 0
		for j < img.Data.Bounds().Max.X {
			r, g, b, a := img.Data.At(j, i).RGBA()
			weightedAverage := (float64(r) * 0.3) + (float64(g) * 0.59) + (float64(b) * 0.11)
			target.Set(j, i, color.NRGBA64{
				R: uint16(weightedAverage),
				G: uint16(weightedAverage),
				B: uint16(weightedAverage),
				A: uint16(a),
			})
			j++
		}
		i++
	}
	img.Data = target
}

func (img *Img) SaveToFile(method string) error {
	outputPath := fmt.Sprintf("parsed_images/%s_method_%s", method, img.Filename)
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	err = jpeg.Encode(file, img.Data, nil)
	if err != nil {
		return err
	}

	return nil
}
