package grabber

import (
	"github.com/lazywei/go-opencv/opencv"
	"log"
)

func GetFrame( capture *opencv.Capture ) *opencv.IplImage {
	if !capture.GrabFrame() {
		log.Println("Not grabe source capture")
	}

	frame := capture.RetrieveFrame(1)

	if frame == nil {
		log.Println("Frame is nil")
	}

	return frame
}
