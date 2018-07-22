package container

import "sort"

//Series
type Series struct {
	data []float64
}

//NewSeries
func NewSeries(l []float64) *Series {
	obj := &Series{
		data: l,
	}
	return obj
}

//Copy return a new Series
func (s *Series) Copy() *Series {
	newS := new(Series)
	// glog.Infoln(len(s.data))
	// glog.Infoln(newS.data)
	newS.data = make([]float64, len(s.data))
	for i, v := range s.data {
		newS.data[i] = v
	}
	return newS
}

//Elem
func (s *Series) Elem(i int) float64 {
	return s.data[i]
}

//Append
func (s *Series) Append(f float64) {
	s.data = append(s.data, f)
}

//ToList
func (s *Series) ToList() []float64 {
	return s.data
}

//Len
func (s *Series) Len() int {
	return len(s.data)
}

//Sort
func (s *Series) Sort() *Series {
	tmp := make([]float64, s.Len())
	copy(tmp, s.data)
	sort.Float64s(tmp)
	return &Series{tmp}
}
