package main

import "fmt"

func pipelineImageParsing(imagePaths []string) {
	fmt.Printf("Images to parse: %d\n", len(imagePaths))
	imagesChan := loadImages(imagePaths)
	resizedChan := resizeImages(imagesChan)
	grayScaledChan := grayScaleImages(resizedChan)
	savedImagesNum := saveImages(grayScaledChan)
	fmt.Printf("Saved images number: %d\n", savedImagesNum)
}

func loadImages(imagePaths []string) <-chan Img {
	out := make(chan Img)
	go func() {
		for _, path := range imagePaths {
			img, err := LoadImage(path)
			if err != nil {
				fmt.Println(err)
				continue
			}
			data := &Img{
				Filename: path,
				Data:     img,
			}
			out <- *data
		}
		close(out)
	}()
	return out
}

func resizeImages(in <-chan Img) <-chan Img {
	out := make(chan Img)
	go func() {
		for data := range in {
			data.Resize()
			out <- data
		}
		close(out)
	}()
	return out
}

func grayScaleImages(in <-chan Img) <-chan Img {
	out := make(chan Img)
	go func() {
		for data := range in {
			data.GrayScale()
			out <- data
		}
		close(out)
	}()
	return out
}

func saveImages(in <-chan Img) int {
	resultChan := make(chan int)
	go func() {
		count := 0
		for data := range in {
			if err := data.SaveToFile("pipeline"); err != nil {
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
