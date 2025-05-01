package main

import (
	"fmt"
	"github.com/jokerinya/image_parser/helpers"
)

func main() {
	imagePaths := []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg"}
	fmt.Printf("Images to parse: %d\n", len(imagePaths))
	imagesChan := loadImages(imagePaths)
	resizedChan := resizeImages(imagesChan)
	grayScaledChan := grayScaleImages(resizedChan)
	savedImagesNum := saveImages(grayScaledChan)
	fmt.Printf("Saved images number: %d\n", savedImagesNum)
}

func loadImages(imagePaths []string) <-chan helpers.Img {
	out := make(chan helpers.Img)
	go func() {
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
			out <- *data
		}
		close(out)
	}()
	return out
}

func resizeImages(in <-chan helpers.Img) <-chan helpers.Img {
	out := make(chan helpers.Img)
	go func() {
		for data := range in {
			data.Resize()
			out <- data
		}
		close(out)
	}()
	return out
}

func grayScaleImages(in <-chan helpers.Img) <-chan helpers.Img {
	out := make(chan helpers.Img)
	go func() {
		for data := range in {
			data.GrayScale()
			out <- data
		}
		close(out)
	}()
	return out
}

func saveImages(in <-chan helpers.Img) int {
	resultChan := make(chan int)
	go func() {
		count := 0
		for data := range in {
			if err := data.SaveToFile(); err != nil {
				fmt.Println(err)
				continue
			}
			count++
		}
		resultChan <- count
		close(resultChan)
	}()
	return <-resultChan
}
