package processing

import (
	"github.com/lazywei/go-opencv/opencv"
)

func ProcessImage(frame *opencv.IplImage, window *opencv.Window) error {
	w := frame.Width()
	h := frame.Height()

	// Create the output image
	cedge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 3)
	defer cedge.Release()

	// Convert to grayscale
	gray := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	edge := opencv.CreateImage(w, h, opencv.IPL_DEPTH_8U, 1)
	defer gray.Release()
	defer edge.Release()

	opencv.CvtColor(frame, gray, opencv.CV_BGR2GRAY)

	opencv.Smooth(gray, edge, opencv.CV_BLUR, 3, 3, 0, 0)
	opencv.Not(gray, edge)

	// Run the edge detector on grayscale
	//opencv.Canny(gray, edge, float64(pos), float64(pos*3), 3)

	opencv.Zero(cedge)
	// copy edge points
	opencv.Copy(frame, cedge, edge)

	window.ShowImage(frame)

	return nil
}
