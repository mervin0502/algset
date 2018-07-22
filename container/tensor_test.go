package container

import "testing"

func TestNewNilTensor(t *testing.T) {
	var ts *Tensor
	ts = NewTensor(nil)
	t.Logf("%s", ts.String())
}
func TestNewTensor(t *testing.T) {

	// var v interface{}
	var ts *Tensor
	//0-d: int
	// v00 := 1
	// ts = NewTensor0(v00)
	// if ts.vtype != VInt {
	// 	t.Errorf("wrong type: %v: %s", v00, ts.vtype)
	// } else {
	// 	t.Logf("input: %v => %s %v", v00, ts.vtype, ts.dimension)
	// }
	// //0-d: uint8

	// //1-d: []float64
	// // var v10 []interface{}
	// v10 := []interface{}{1.1, 2.2, 3.3, 4.4, 5.5}
	// ts = NewTensor1(v10)
	// if ts.vtype != VFloat64 {
	// 	t.Errorf("wrong type: %v: %s", v10, ts.vtype)
	// } else {
	// 	t.Logf("input: %v => %s %v", v10, ts.vtype, ts.dimension)
	// }

	//NewTensor
	// vv00 := 1
	// ts = NewTensor(vv00)
	// t.Logf("%s", ts)
	// vv01 := "ssss"
	// ts = NewTensor(vv01)
	// t.Logf("%s", ts)
	// vv02 := false
	// ts = NewTensor(vv02)
	// t.Logf("%s", ts)
	// vv03 := 9.333
	// ts = NewTensor(vv03)
	// t.Logf("%s", ts)

	// //1-d
	// vv10 := []int{1, 2, 3}
	// ts = NewTensor(vv10)
	// t.Logf("%s", ts)
	// vv11 := []float64{1.1, 1.2, 1.3}
	// ts = NewTensor(vv11)
	// t.Logf("%s", ts)
	// //2-d
	// vv20 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}}
	// ts = NewTensor(vv20)
	// t.Logf("%s", ts)
	vv21 := []interface{}{[]int{1, 2, 3}, []int{4, 5, 6}}
	ts = NewTensor(vv21)

	t.Logf("%s", ts.String())

	vv22 := []interface{}{[]int{1, 2, 3}, []float64{4.0, 5.5, 6.9}, []string{"a", "b", "c"}}
	ts = NewTensor(vv22)

	t.Logf("%s", ts.String())

	var ts1 *Tensor
	vv22 = []interface{}{[]int{1, 2, 3, 3}, []float64{4.0, 5.5, 6, 9}, []string{"a", "b", "c", "c"}}
	ts = NewTensor(vv22)
	_d, _l := ts.Shape()
	t.Logf("%d, %v", _d, _l)

	ts1 = NewTensor(ts)

	t.Logf("%s", ts1)
}

func TestDimension(t *testing.T) {
	var ts *Tensor
	vv22 := []interface{}{[]int{1, 2, 3}, []float64{4.0, 5.5, 2}, []string{"a", "b", "c"}}
	ts = NewTensor(vv22)

	t.Logf("%s", ts.String())
}

func TestShape(t *testing.T) {
	var ts *Tensor
	vv22 := []interface{}{[]int{1, 2, 3, 3}, []float64{4.0, 5.5, 6, 9}, []string{"a", "b", "c", "c"}}
	ts = NewTensor(vv22)
	_d, _l := ts.Shape()
	t.Logf("%d, %v", _d, _l)

	vv01 := 0
	ts = NewTensor(vv01)
	_d, _l = ts.Shape()
	t.Logf("%d, %v", _d, _l)

	vv11 := []int{2, 4, 0, 4, 5, 8}
	ts = NewTensor(vv11)
	_d, _l = ts.Shape()
	t.Logf("%d, %v", _d, _l)
}

func TestEqual(t *testing.T) {
	var ts1, ts2 *Tensor
	vv220 := []interface{}{[]int{1, 2, 3, 3}, []float64{4.0, 5.5, 6, 9}, []string{"a", "b", "c", "c"}}
	ts1 = NewTensor(vv220)
	vv221 := []interface{}{[]int{1, 2, 3, 3}, []float64{4.0, 5.5, 6, 9}, []string{"a", "b", "c", "c"}}
	ts2 = NewTensor(vv221)
	t.Logf("%v", ts1.Equal(ts2))

	vv222 := []interface{}{[]int{1, 2, 3, 3}, []float64{4.0, 5.5, 6, 9.1}, []string{"a", "b", "c", "c"}}
	ts2 = NewTensor(vv222)
	t.Logf("%v", ts1.Equal(ts2))
}
