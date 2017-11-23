package output

import (
	"github.com/lazywei/go-opencv/opencv"
)

func GetWindow() *opencv.Window {
	window := opencv.NewWindow("Video")
	return window
}
