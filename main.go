package main

import (
	"fmt"
	"time"
)

var imagePaths = []string{"image1.jpg", "image2.jpg", "image3.jpg", "image4.jpg"}

func main() {
	run(sequentialImageParsing, sequential) // ~210-220 ms.
	fmt.Println("----------")
	run(pipelineImageParsing, pipeline) // ~180-190 ms.
	fmt.Println("----------")
	run(fanInFanOutImageParsing, fanInFanOut) // ~70 ms.
}

func run(parser func([]string), parserName string) {
	start := time.Now()
	parser(imagePaths)
	elapsed := time.Since(start)
	fmt.Printf("%s elapsed: %s\n", parserName, elapsed)
}
