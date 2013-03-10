package opencv

import (
	"fmt"
	"testing"
)

func TestCreateHist(t *testing.T) {
	ranges := []float64{0, 255}
	h := CreateHist(1, 256, CV_HIST_ARRAY, ranges, true)
	defer h.Release()
	h.Clear()
}

func TestCalcHist(t *testing.T) {
	img := LoadImage("./iguana.jpg")

	if img == nil {
		t.Error("Could not load image")
	}
	grayImg := CreateImage(img.Width(),
		img.Height(), IPL_DEPTH_8U, 1)

	CvtColor(img, grayImg, CV_BGR2GRAY)

	ranges := []float64{0, 255}

	h := CreateHist(1, 256, CV_HIST_ARRAY, ranges, true)
	defer h.Release()

	images := []*IplImage{grayImg}
	h.CalcHist(images, false, nil)
	fmt.Printf("%v", h)
	defer h.Clear()
}
