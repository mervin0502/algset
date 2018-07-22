package statistic

import (
	"github.com/golang/glog"
)

//MeanOfFloat return the average of the array
func MeanOfFloat(data []float64) float64 {

	if len(data) == 0 {
		glog.Fatalln("the data is empty.")
	}
	var sum float64
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}
