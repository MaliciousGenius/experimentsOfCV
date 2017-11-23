package main

import (
	capture "./capture"
	output "./output"
	grabber "./grabber"
	processing "./processing"
	"github.com/lazywei/go-opencv/opencv"
	"os"
)

func main() {
	sourceCapture := capture.GetIPcam()
	defer sourceCapture.Release()

	window := output.GetWindow()
	defer window.Destroy()

	for true {
		frame := grabber.GetFrame(sourceCapture)

		processing.ProcessImage(frame, window)

		if key := opencv.WaitKey(10); key == 27 {
			os.Exit(0)
		}
	}
	opencv.WaitKey(0)
}
