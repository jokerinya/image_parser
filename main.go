package main

import (
	"fmt"
	"github.com/jokerinya/image_parser/helpers"
)

func main() {
	imagePaths := []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg"}
	for _, path := range imagePaths {
		img, err := helpers.LoadImage(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data := &helpers.Img{
			Filename: path,
			Data:     img,
		}
		data.Resize()
		data.GrayScale()
		err = data.SaveToFile()
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
