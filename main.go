package main

import (
	capture "./capture"
	display "./display"
)

func main() {
	sourceCapture := capture.GetWebcam()
	defer sourceCapture.Release()
	display.ShowVideo(sourceCapture)
}
