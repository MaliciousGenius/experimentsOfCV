package capture

import (
	"github.com/lazywei/go-opencv/opencv"
)

func GetIPcam() *opencv.Capture {
	cameraCapture := opencv.NewFileCapture("http://71.2.17.14/mjpg/video.mjpg")

	if cameraCapture == nil {
		panic("ERROR: can not open IP camera")
	}

	return cameraCapture
}
