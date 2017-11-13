package main

import (
	capture "github.com/experimentsOfCV/capture"
	display "github.com/experimentsOfCV/display"
)

func main() {
	sourceCapture := capture.GetWebcam()
	defer sourceCapture.Release()
	display.ShowVideo(sourceCapture)
}
