package main

import (

)

func main() {
	sourceCapture := GetWebcam()
	defer sourceCapture.Release()
	ShowVideo(sourceCapture)
}
