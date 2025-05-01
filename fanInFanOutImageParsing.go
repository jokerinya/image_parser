package main

import (
	"fmt"
	"sync"
)

// Almost fan-in, fan-out :)

const fanInFanOut = "fanInFanOut"

func fanInFanOutImageParsing(imagePaths []string) {
	fmt.Printf("Images to parse: %d\n", len(imagePaths))
	var wg sync.WaitGroup

	res := make(chan bool)
	for _, imagePath := range imagePaths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			err := individualImageParsing(path)
			if err != nil {
				fmt.Println(err)
				return
			}
			res <- true
		}(imagePath)
	}

	// another goroutine to close the channel
	go func() {
		wg.Wait()
		close(res)
	}()

	out := 0
	for range res {
		out++
	}
	fmt.Printf("Saved images number: %d\n", out)
}

func individualImageParsing(path string) error {
	img, err := LoadImage(path)
	if err != nil {
		return err
	}
	data := &Img{
		Data:     img,
		Filename: path,
	}
	data.Resize()
	data.GrayScale()
	err = data.SaveToFile(fanInFanOut)
	if err != nil {
		return err
	}
	return nil
}
