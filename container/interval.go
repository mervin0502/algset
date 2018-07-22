package container

import "fmt"

//Interval
type Interval struct {
	Start float64
	End   float64
}

//NewInterval return the new Interval struct
func NewInterval(s, e float64) *Interval {
	i := &Interval{
		Start: s,
		End:   e,
	}
	return i
}

//String
func (i *Interval) String() string {
	return fmt.Sprintf("[%.6f, %.6f]", i.Start, i.End)
}
