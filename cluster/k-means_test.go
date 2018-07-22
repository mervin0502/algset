package cluster

import (
	"testing"

	"mervin.me/algset/container"
)

func TestInitSeriesCenter(t *testing.T) {

	l1 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	s1 := container.NewTensor(l1)
	k := 3
	c1, err := randomizeTensorCenteroid(s1, k)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", c1)
}

func TestInitTensorCenter2d(t *testing.T) {
	l1 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l2 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l3 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l4 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l5 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l6 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l7 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l8 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l9 := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	l := [][]float64{l1, l2, l3, l4, l5, l6, l7, l8, l9}
	s := container.NewTensor(l)
	k := 3
	c1, err := randomizeTensorCenteroid(s, k)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", c1)
}

func Test2DPartion(t *testing.T) {
	data := []interface{}{
		[]float64{20.0, 20.0, 20.0, 20.0},
		[]float64{21.0, 21.0, 21.0, 21.0},
		[]float64{100.5, 100.5, 100.5, 100.5},
		[]float64{50.1, 50.1, 50.1, 50.1},
		[]float64{64.2, 64.2, 64.2, 64.2},
	}
	ts := container.NewTensor(data)
	_g, _c := KMeans(ts, 3, 50)
	t.Logf("%v \n\n %v", _g, _c)
}

func Test1DPartion(t *testing.T) {
	data := []float64{2, 4, 5, 3, 4, 8, 9, 10, 8, 11, 20, 21, 23, 22, 21}
	ts := container.NewTensor(data)
	_g, _c := KMeans(ts, 3, 50)
	t.Logf("%v \n\n %v", _g, _c.SliceString())
}
