package container

import (
	"errors"
	"flag"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

func init() {
	flag.Parse()
	flag.Set("logtostderr", "true")
}

/*
0 - variable(size) - int, float, string, boolean, complex
1 - vector(size and direction) - v = [1, 2, 3]
2 - matrix(data sheet) - m = [[1, 2, 3], [4, 5, 6], [7,8,9]]
3 - 3-dimension - t = [[[1], [2], [3]], [[4], [5], [6]],[[7], [8], [9]],]
x - x-dimension
*/
//TensorDimension
// type TensorDimension int

// const (
// 	Tensor0D TensorDimension = iota
// 	Tensor1D
// 	Tensor2D
// 	Tensor3D
// 	TensorXD
// )

//Tensor define the Tensor struct
type Tensor struct {
	Label     string
	Value     interface{}
	Dimension int
	VType     reflect.Kind
	Length    int
}

//NewTensor create a new Tensor
func NewTensor(v interface{}) *Tensor {

	var t Tensor
	//vt -> nil
	if v == nil {
		// glog.Infof("%v", v)
		return &Tensor{
			Value:     nil,
			Dimension: -1,
			Length:    -1,
			VType:     reflect.Invalid,
		}
	}
	vt := reflect.TypeOf(v)
	//v -> tensor
	if strings.Contains(vt.String(), "Tensor") {
		v0, ok := v.(*Tensor)
		if ok {
			t.Dimension = v0.Dimension + 1
			t.Length = 1
			t.VType = reflect.Slice
			t.Label = v0.Label
			t.Value = []*Tensor{v0}
			return &t
		}
		v1, ok := v.(Tensor)
		if ok {
			t.Dimension = v1.Dimension + 1
			t.Length = 1
			t.VType = reflect.Slice
			t.Label = v1.Label
			t.Value = []*Tensor{&v1}
			return &t
		}
	}
	_d, _l, _vt := parserTensor(v, &t)
	t.Dimension = _d
	t.Length = _l
	t.VType = _vt
	return &t
}

//parserTensor
func parserTensor(v interface{}, t *Tensor) (depth, lenPerDim int, k reflect.Kind) {
	vt := reflect.TypeOf(v)
	vv := reflect.ValueOf(v)

	// glog.Infof("%v %s", v, vt)
	switch vt.Kind() {
	case reflect.Bool:
		fallthrough
	case reflect.String:
		fallthrough
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		fallthrough
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		fallthrough
	case reflect.Float32, reflect.Float64:
		fallthrough
	case reflect.Complex64, reflect.Complex128:
		fallthrough
	case reflect.Struct:
		if strings.Contains(vt.String(), "Tensor") {
			v0, ok := v.(Tensor)
			if ok {
				t.Value = &v0.Value
				t.Label = v0.Label
				return v0.Dimension, v0.Length, v0.VType
			}
		}
		t.Value = v
		return 0, 1, vt.Kind()
	case reflect.Ptr:
		if strings.Contains(vt.String(), "Tensor") {
			v0, ok := v.(*Tensor)
			if ok {
				t.Value = v0.Value
				t.Label = v0.Label
				return v0.Dimension, v0.Length, v0.VType
			}
		}
		t.Value = v
		return 0, 1, vt.Kind()
	case reflect.Slice:
		_data := make([]*Tensor, vv.Len())
		var _d, _dd, _l, _ll int
		var _vt reflect.Kind
		for i := 0; i < vv.Len(); i++ {
			var tt Tensor
			_d, _l, _vt = parserTensor(vv.Index(i).Interface(), &tt)
			if i > 0 {
				if _l != _ll {
					glog.Fatalf("the Length of %dth is wrong.", i)
				}
				if _d != _dd {
					glog.Fatalf("the dimension of %dth is wrong.", i)
				}
			}
			_dd = _d
			_ll = _l
			tt.Dimension = _d
			tt.Length = _l
			tt.VType = _vt
			_data[i] = &tt
		}
		t.Value = _data
		return _dd + 1, vv.Len(), vt.Kind()
	}
	return -1, -1, vt.Kind()
}

//Put put a sub-value into the t.Value with the index i
func (t *Tensor) Put(i int, tt *Tensor) bool {
	if t.Length == 1 {
		glog.Errorln("the value of the t is not an array")
		return false
	}
	if i > t.Length {
		glog.Errorln("the i is out of the t.Value")
		return false
	}
	vs, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Errorln(ErrTypeAssertion)
		return false
	}
	if !vs[i].DimEqual(tt) {
		glog.Errorf("%v \n %v", vs[i], tt)
		glog.Errorln(ErrDimensionNotEqual)
		return false
	}
	vs[i] = tt
	t.Value = vs
	return true
}

//Append append the v into the t.Value
func (t *Tensor) Append(v interface{}) *Tensor {
	/*
	   1 <- 0
	   2 <-1

	*/
	var subT Tensor
	_d, _l, _vt := parserTensor(v, &subT)
	subT.Dimension = _d
	subT.Length = _l
	subT.VType = _vt

	//blank tensor: -1
	if t.Dimension == -1 {
		t.Value = []*Tensor{&subT}
		t.Dimension = subT.Dimension + 1
		t.Length = 1
		t.VType = reflect.Slice
		return &subT
	}
	//0-dimension
	if t.Dimension == 0 {
		glog.Errorf("%s", "can not append to the 0-Dimension")
		return nil
	}

	if t.Dimension-subT.Dimension != 1 {
		glog.Errorf("%s", "can not append to the t")
		return nil
	}
	// glog.Infof("%v", t.VType)
	if t.Length == 1 && t.VType != reflect.Slice {
		tt, ok := t.Value.(*Tensor)
		if !ok {
			glog.Errorln(ErrTypeAssertion)
			return nil
		}
		_data := make([]*Tensor, 2)
		if !subT.DimEqual(tt) {
			glog.Errorln(ErrDimensionNotEqual)
			return nil
		}
		_data[0] = tt
		_data[1] = &subT
		t.Value = _data
		t.Length = 2
	} else {
		tt, ok := t.Value.([]*Tensor)
		if !ok {
			glog.Errorln(ErrTypeAssertion)
			return nil
		}
		if !subT.DimEqual(tt[0]) {
			glog.Errorln(ErrDimensionNotEqual)
			return nil
		}
		tt = append(tt, &subT)
		t.Value = tt
		t.Length = t.Length + 1
	}

	return &subT
}

//Tanspose return a transposed tensor
func (t *Tensor) Tanspose() *Tensor {
	if t.Dimension != 2 {
		return t
	}
	newTensor := new(Tensor)
	vs, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Fatalln(ErrTypeAssertion)
	}
	var tmpTensor *Tensor
	var tss [][]*Tensor
	tss = make([][]*Tensor, vs[0].Length)
	for i, v := range vs {
		vs1, ok := v.Value.([]*Tensor)
		if !ok {
			glog.Fatalln(ErrTypeAssertion)
		}
		for i1, v1 := range vs1 {
			if tss[i1] == nil {
				tss[i1] = make([]*Tensor, t.Length)
			}
			tss[i1][i] = v1.Copy()
		}
	} //for
	var ts []*Tensor
	ts = make([]*Tensor, vs[0].Length)
	for i, v := range tss {
		tmpTensor = new(Tensor)
		tmpTensor.Dimension = t.Dimension - 1
		tmpTensor.Length = t.Length
		tmpTensor.VType = reflect.Slice
		tmpTensor.Label = strconv.Itoa(i)
		tmpTensor.Value = v
		ts[i] = tmpTensor
	}

	newTensor.Dimension = t.Dimension
	newTensor.Length = vs[0].Length
	newTensor.VType = t.VType
	newTensor.Label = t.Label
	newTensor.Value = ts
	return newTensor
}

//Copy return a new tensor
func (t *Tensor) Copy() *Tensor {

	newT := new(Tensor)
	if t.Dimension == 0 {
		newT.Value = t.Value
		newT.Dimension = t.Dimension
		newT.Length = t.Length
		newT.VType = t.VType
		newT.Label = t.Label
		return newT
	}
	tData, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Errorln(ErrTypeAssertion)
		return nil
	}
	_data := make([]*Tensor, t.Length)
	for i, tt := range tData {
		_data[i] = tt.Copy()
	}
	newT.Value = _data
	newT.Dimension = t.Dimension
	newT.Length = t.Length
	newT.VType = t.VType
	newT.Label = t.Label
	return newT
}

//Get return the ith tensor from 0 ~ n-1
func (t Tensor) Get(i int) *Tensor {
	if i >= t.Length {
		return nil
	}
	if t.Dimension == -1 {
		return nil
	}
	if t.Dimension == 0 {
		return &t
	} else {
		v, ok := t.Value.([]*Tensor)
		if !ok {
			glog.Fatalln("wrong VType")
		}
		// glog.Infof("%s", t.String())
		// glog.Infof("%d - %d -  %d", i, t.Length, len(v))
		return v[i]
	}

}

//Filter filter the t which include the row index and column index
func (t *Tensor) Filter(vIndex, subVIndex []int) *Tensor {
	maxFunc := func(is []int) int {
		max := 0
		for _, v := range is {
			if v > max {
				max = v
			}
		}
		return max
	}
	containFunc := func(is []int, i int) bool {
		for _, v := range is {
			if v == i {
				return true
			}
		}
		return false
	}
	maxRIndex := maxFunc(vIndex)
	maxCIndex := maxFunc(subVIndex)

	vs, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Fatalln(ErrTypeAssertion)
		return nil
	}
	if maxRIndex > t.Length {
		glog.Fatalln(errors.New("the maximum of vIndex is out of the t.Length"))
		return nil
	}
	if maxCIndex > vs[0].Length {
		glog.Fatalln(errors.New("the maximum of subVIndex is out of the t[0].Length"))
		return nil
	}
	var tt *Tensor
	var sub *Tensor
	var ts, ts1 []*Tensor
	tt = new(Tensor)
	tt.Dimension = t.Dimension
	tt.Length = len(vIndex)
	tt.VType = t.VType
	tt.Label = t.Label
	ts = make([]*Tensor, len(vIndex))

	var ii, jj int
	for i, v := range vs {
		if !containFunc(vIndex, i) {
			continue
		}
		vs1, ok := v.Value.([]*Tensor)
		if !ok {
			glog.Fatalln(ErrTypeAssertion)
			return nil
		}
		ts1 = make([]*Tensor, len(subVIndex))
		jj = 0
		for j, v1 := range vs1 {
			if !containFunc(subVIndex, j) {
				continue
			}
			ts1[jj] = v1
			jj++
		}
		sub = new(Tensor)
		sub.Dimension = v.Dimension
		sub.Length = len(subVIndex)
		sub.VType = v.VType
		sub.Label = v.Label
		sub.Value = ts1
		ts[ii] = sub
		ii++
	}
	tt.Value = ts
	return tt
}

//Values return the array t.Value
func (t *Tensor) Values() []*Tensor {
	if t.Dimension == -1 {
		return nil
	}
	if t.Dimension == 0 {
		return []*Tensor{t}
	}
	v, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Errorln(ErrTypeAssertion)
	}
	return v
}

//Shape return the diemension and the length of each of dimensions
func (t Tensor) Shape() (int, []int) {

	if t.Dimension == 0 {
		return 0, []int{}
	}

	v, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Errorln(ErrTypeAssertion)
	}
	_, _ll := v[0].Shape()
	out := append([]int{t.Length}, _ll...)
	return t.Dimension, out
}

//Empty return the true/false of the Tensor
func (t Tensor) Empty() bool {
	if t.Dimension == 0 {
		if t.Length == 0 {
			return true
		} else {
			return false
		}
	}
	return false
}

//Equal return true if the t == tt
func (t *Tensor) Equal(tt *Tensor) bool {

	return reflect.DeepEqual(t, tt)
}

//DimEqual return true if the dimensions of t and tt are equal
func (t *Tensor) DimEqual(tt *Tensor) bool {
	if t.Dimension != tt.Dimension {
		return false
	}
	if t.Length != tt.Length {
		return false
	}
	if t.Dimension == 0 || t.Dimension == -1 {
		return true
	}
	tData, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Errorln(ErrTypeAssertion)
	}
	ttData, ok := tt.Value.([]*Tensor)
	if !ok {
		glog.Errorln(ErrTypeAssertion)
	}
	return tData[0].DimEqual(ttData[0])
}

/*

format output
*/

//Int return the int value
func (t Tensor) Int() (int, error) {
	if t.Dimension != 0 {
		return -1, errors.New("the t's dimension is greater than 0")
	}
	if t.Length != 1 {
		return -1, errors.New("the t.Value is not a simple value")
	}
	v, ok := t.Value.(int)
	if !ok {
		return -1, ErrTypeAssertion
	}
	return v, nil
}

//Int64 return the int64 value
func (t Tensor) Int64() (int64, error) {
	if t.Dimension != 0 {
		return -1, errors.New("the t's dimension is greater than 0")
	}
	if t.Length != 1 {
		return -1, errors.New("the t.Value is not a simple value")
	}
	v, ok := t.Value.(int64)
	if !ok {
		return -1, ErrTypeAssertion
	}
	return v, nil
}

//Float64 return the float64 value
func (t Tensor) Float64() (float64, error) {
	if t.Dimension != 0 {
		return -1, errors.New("the t's dimension is greater than 0")
	}
	if t.Length != 1 {
		return -1, errors.New("the t.Value is not a simple value")
	}
	v, ok := t.Value.(float64)
	if !ok {
		return -1, ErrTypeAssertion
	}
	return v, nil
}

//String return the format output of the t
func (t Tensor) String() string {

	// glog.Infof("{dim=%d, len=%d, VType=%s}", t.Dimension, t.Length, t.VType)
	if t.Dimension == -1 {
		return fmt.Sprintf(`{"dim":%d, "len":%d, "vType":"%s", "label":"%s", "value":%v}`, t.Dimension, t.Length, t.VType, t.Label, t.Value)
	}
	if t.Dimension == 0 {
		return fmt.Sprintf(`{"dim":%d, "len":%d, "vType":"%s", "label":"%s", "value":%v}`, t.Dimension, t.Length, t.VType, t.Label, t.Value)
	}
	// glog.Infoln(reflect.TypeOf(t.Value).Kind())
	// glog.Infoln("####")
	// glog.Infof("%s", t.Value)
	v, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Fatalln("wrong type")
	}
	vvStr := "["
	for _, vv := range v {
		vvStr += vv.String() + ","
	}
	vvStr = vvStr[0:len(vvStr)-1] + "]"
	return fmt.Sprintf(`{"dim":%d, "len":%d, "vType":"%s", "label":"%s", "value":%v}`, t.Dimension, t.Length, t.VType, t.Label, vvStr)

}

//SliceString return the slice format output of the t
func (t Tensor) SliceString() string {
	// glog.Infof("{dim=%d, len=%d, VType=%s}", t.Dimension, t.Length, t.VType)
	if t.Dimension == -1 {
		return "nil"
	}
	if t.Dimension == 0 {
		return fmt.Sprintf("%v", t.Value)
	}
	v, ok := t.Value.([]*Tensor)
	if !ok {
		glog.Fatalln("wrong type")
	}
	vvStr := "["
	for _, vv := range v {
		vvStr += vv.SliceString() + ","
	}
	vvStr = vvStr[0:len(vvStr)-1] + "]"
	return vvStr
}
