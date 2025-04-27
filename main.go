package main

import (
	"fmt"
	"github.com/jokerinya/image_parser/helpers"
)

func main() {
	imagePaths := []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg"}
	for _, path := range imagePaths {
		_, err := helpers.LoadImage(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("")
	}
}
