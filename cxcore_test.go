package opencv

import (
	"testing"
)

func TestLoadImage2(t *testing.T) {
	// t.Errorf("aaa")
}

func TestCreateHist(t *testing.T) {
	img := LoadImage("./iguana.jpg")

	if img == nil {
		t.Error("Could not load image")
	}

	ranges := []float64{0, 255}
	h := CreateHist(1, 256, CV_HIST_ARRAY, ranges, true)
	defer h.Release()
	h.Clear()
}
