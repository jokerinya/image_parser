package main

import (
	"fmt"
	"github.com/jokerinya/image_parser/helpers"
)

func main() {
	imagePaths := []string{
		"raw_images/image1.jpg",
		"raw_images/image2.jpg",
		"raw_images/image3.jpg",
		"raw_images/image4.jpg",
	}
	for _, path := range imagePaths {
		_, err := helpers.LoadImage(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("")
	}
}
