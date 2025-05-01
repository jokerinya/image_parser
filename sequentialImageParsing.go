package main

import (
	"fmt"
	"github.com/jokerinya/image_parser/helpers"
)

func sequentialImageParsing(imagePaths []string) {
	fmt.Printf("Images to parse: %d\n", len(imagePaths))
	out := 0
	for _, path := range imagePaths {
		img, err := helpers.LoadImage(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data := &helpers.Img{
			Data:     img,
			Filename: path,
		}
		data.Resize()
		data.GrayScale()
		err = data.SaveToFile()
		if err != nil {
			fmt.Println(err)
			continue
		}
		out++
	}
	fmt.Printf("Saved images number: %d\n", out)
}
