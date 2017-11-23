package grabber

import (
	"github.com/lazywei/go-opencv/opencv"
	"log"
)

func GetFrame( capture *opencv.Capture ) *opencv.IplImage {
	var frame *opencv.IplImage

	if capture.GrabFrame() {
		data := capture.RetrieveFrame(1)
		if data == nil {
			log.Println("Frame is nil")
		}
	frame = data.Clone()
	} else {
		log.Println("Not grabe source capture")
	}

	return frame
}
