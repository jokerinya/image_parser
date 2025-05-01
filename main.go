package main

import (
	"fmt"
	"time"
)

func main() {
	imagePaths := []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg"}
	runSequentialImageParsing(imagePaths) // ~210-220 ms.
	fmt.Println("----------")
	runAsyncPipelineImageParsing(imagePaths) // ~180-190 ms.
}

func runSequentialImageParsing(imagePaths []string) {
	start := time.Now()
	sequentialImageParsing(imagePaths)
	elapsed := time.Since(start)
	fmt.Println("sequential elapsed: ", elapsed)
}

func runAsyncPipelineImageParsing(imagePaths []string) {
	start := time.Now()
	pipelineImageParsing(imagePaths)
	elapsed := time.Since(start)
	fmt.Println("async elapsed: ", elapsed)
}
