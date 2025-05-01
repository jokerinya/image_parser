package main

import (
	"fmt"
	"time"
)

func main() {
	imagePaths := []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg"}
	start := time.Now()
	pipelineImageParsing(imagePaths)
	elapsed := time.Since(start)
	fmt.Println("async elapsed: ", elapsed)
}
