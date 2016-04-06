package main

import (
	"fmt"
	"os"
	"time"

	"github.com/lazywei/go-opencv/opencv"
)

func main() {
	filename := os.Args[1]

	cap := opencv.NewCameraCapture(0)
	if cap == nil {
		panic("No webcam found")
	}
	defer cap.Release()

	// let it warm up for 10s - autofocus and w/e
	for i := 0; i < 10; i++ {
		cap.QueryFrame()
		fmt.Print(".")
		time.Sleep(time.Second)
	}
	fmt.Println("!")
	// actually take the picture
	img := cap.QueryFrame()
	if img == nil {
		panic("The image was null?!?")
	}
	opencv.SaveImage(filename, img, 0)
}
