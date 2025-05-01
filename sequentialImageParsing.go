package main

import "fmt"

const sequential = "sequential"

func sequentialImageParsing(imagePaths []string) {
	fmt.Printf("Images to parse: %d\n", len(imagePaths))
	out := 0
	for _, path := range imagePaths {
		img, err := LoadImage(path)
		if err != nil {
			fmt.Println(err)
			continue
		}
		data := &Img{
			Data:     img,
			Filename: path,
		}
		data.Resize()
		data.GrayScale()
		err = data.SaveToFile(sequential)
		if err != nil {
			fmt.Println(err)
			continue
		}
		out++
	}
	fmt.Printf("Saved images number: %d\n", out)
}
