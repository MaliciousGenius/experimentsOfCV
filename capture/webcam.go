package capture

import (
	"github.com/lazywei/go-opencv/opencv"
)

func GetWebcam() *opencv.Capture {
	cameraCapture := opencv.NewCameraCapture(0)

	if cameraCapture == nil {
		panic("ERROR: can not open camera")
	}

	return cameraCapture
}
