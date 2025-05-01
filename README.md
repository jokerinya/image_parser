# Image parser
# Image Processing Pipeline (Go)

This project demonstrates image processing using both **sequential** and **asynchronous pipeline** approaches in Go.

## Features

- Loads images (`loadImages`)
- Resizes them (`resizeImages`)
- Converts to grayscale (`grayScaleImages`)
- Saves them to disk (`saveImages`)
- Measures and prints execution time

## Usage

```bash
go run main.go
```

### Sample Output

```
Images to parse: 4
Saved images number: 4
sequential elapsed:  215.328125ms
----------
Images to parse: 4
Saved images number: 4
async elapsed:  184.294792ms
```

## Purpose

The goal is to showcase how Go's concurrency primitives — **goroutines** and **channels** — can be used to build a performant data processing pipeline.

## Requirements

- Go 1.18 or higher
- Images: `image1.jpg`, `image2.jpg`, `image3.jpg`, `image4.jpg` must be present in the project `raw_images` folder
- A `helpers` package with the following methods:
    - `LoadImage(path string) (image.Image, error)`
    - `(*Img) Resize()`
    - `(*Img) GrayScale()`
    - `(*Img) SaveToFile() error`

## License

MIT
