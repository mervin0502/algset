package distance

import (
	"math"

	"mervin.me/algset/container"
)

//https://blog.csdn.net/taoyanqi8932/article/details/53727841
//https://blog.csdn.net/google19890102/article/details/26149927
func Euclidean(x, y *container.Tensor) (float64, error) {
	if x.Dimension != y.Dimension {
		return -1, container.ErrDimensionNotEqual
	}
	if x.Dimension == 0 {
		vx, ok := x.Value.(float64)
		if !ok {
			return -1, container.ErrTypeAssertion
		}
		vy, ok := y.Value.(float64)
		if !ok {
			return -1, container.ErrTypeAssertion
		}
		return math.Abs(vx - vy), nil
	} else if x.Dimension == 1 {
		if x.Length != y.Length {
			return -1, container.ErrLengthNotEqual
		}
		vx := x.Values()
		vy := y.Values()
		var sum float64 = 0.0
		for i := 0; i < x.Length; i += 1 {
			vvx, ok := vx[i].Value.(float64)
			if !ok {
				return -1, container.ErrTypeAssertion
			}
			vvy, ok := vy[i].Value.(float64)
			if !ok {
				return -1, container.ErrTypeAssertion
			}
			// glog.Infof("%f \t %f ", vvx, vvy)
			sum += math.Pow((vvx - vvy), 2)
		}
		return math.Sqrt(sum), nil
	}
	return -1, container.ErrDimensionTooHigh
}
