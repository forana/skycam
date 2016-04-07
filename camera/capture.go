package camera

import (
	"fmt"
	"os"
	"time"

	"github.com/lazywei/go-opencv/opencv"
)

/*
Capture saves a picture from the specified camera to the specified output directory.

The image will be saved in the format `YYYYMMDD-HHMMSS-(timestamp).jpg`.
If present, latest.jpg will also be replaced with that same filename.
*/
func Capture(outputDir string, cameraIndex int) {
	filename := fmt.Sprintf("%s%c%s.jpg", outputDir, os.PathSeparator, ftime())
	latestFilename := fmt.Sprintf("%s%clatest.jpg", outputDir, os.PathSeparator)

	cap := opencv.NewCameraCapture(opencv.CV_CAP_ANY + cameraIndex)
	if cap == nil {
		panic("No webcam found")
	}
	defer cap.Release()

	// let it warm up for 10s - autofocus and w/e
	for i := 0; i < 10; i++ {
		_ = cap.QueryFrame()
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
	opencv.SaveImage(latestFilename, img, 0)

	fmt.Println("Saved to " + filename)
}

func ftime() string {
	t := time.Now()
	return fmt.Sprintf("%04d%02d%02d-%02d%02d%02d-%d", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Unix())
}
