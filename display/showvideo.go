package display

import (
	"log"
	"os"
	"github.com/lazywei/go-opencv/opencv"
)

func ShowVideo(capture *opencv.Capture) error {
	window := opencv.NewWindow("Video")
	defer window.Destroy()

	for {
		if capture.GrabFrame() {
			frame := capture.RetrieveFrame(1)
			if frame == nil {
				log.Println("Frame is nil")
			}
			window.ShowImage(frame)
		} else {
			log.Println("Not grabe source capture")
		}
		if key := opencv.WaitKey(10); key == 27 {
			os.Exit(0)
		}
	}
	opencv.WaitKey(0)
	return nil
}
