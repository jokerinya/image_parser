package main

import (
	"fmt"
	"time"
)

func main() {
	imagePaths := []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg"}
	start := time.Now()
	sequentialImageParsing(imagePaths)
	elapsed := time.Since(start)
	fmt.Println("sequential elapsed: ", elapsed)

	fmt.Println("----------")

	start = time.Now()
	pipelineImageParsing(imagePaths)
	elapsed = time.Since(start)
	fmt.Println("async elapsed: ", elapsed)
}
