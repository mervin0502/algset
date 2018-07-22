package statistic

import "math"

//StdOfFloat return the standard variate of the array
func StdOfFloat(data []float64) float64 {
	avg := MeanOfFloat(data)
	var out float64
	for _, v := range data {
		out += math.Pow(v-avg, 2)
	}
	return math.Sqrt(out / float64(len(data)))
}
