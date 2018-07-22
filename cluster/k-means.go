package cluster

import (
	"errors"
	"flag"
	"math/rand"
	"time"

	"github.com/golang/glog"
	"mervin.me/algset/container"
	"mervin.me/algset/distance"
)

func init() {
	flag.Parse()

	flag.Set("alsologtostderr", "true")
	flag.Set("v", "2")
}

//KMeans xxx
//https://blog.csdn.net/google19890102/article/details/26149927
//在基本k-means的基础上发展而来二分 (bisecting) k-means，其主要思想：一个大cluster进行分裂后可以得到两个小的cluster；为了得到k个cluster，可进行k-1次分裂。
//https://github.com/bugra/kmeans/blob/master/kmeans.go
//https://github.com/muesli/kmeans/blob/master/kmeans.go
//https://github.com/mash/gokmeans/blob/master/gokmeans.go
func KMeans(s *container.Tensor, k, iter int) (map[int]*container.Tensor, *container.Tensor) {

	// var c []float64
	// var f float64
	// s1 := s.Copy()
	// slen := s.Len()
	c, _ := randomizeTensorCenteroid(s, k)
	// glog.Infof("%v", c)

	return trainOfKMeans(s, c, k, iter)
}

//randomizeTensorCenteroid
func randomizeTensorCenteroid(t *container.Tensor, k int) (*container.Tensor, error) {

	if t.Length == 0 {
		return nil, errors.New("the t is null")
	}
	if k == 0 {
		return nil, errors.New("k must be greater than 0")
	}
	var tt *container.Tensor
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	v := t.Values()
	// glog.Infof("%s", t)
	for i := 0; i < k; i++ {
		srcIndex := r.Intn(t.Length)
		// glog.Infof("rand index: %d", srcIndex)
		// if
		if tt == nil {
			tt = container.NewTensor(v[srcIndex].Copy())
		} else {
			tt.Append(v[srcIndex].Copy())
		}
	}
	return tt, nil
}

//trainOfKMeans
func trainOfKMeans(s, c *container.Tensor, k, iter int) (map[int]*container.Tensor, *container.Tensor) {
	// var t *container.Tensor
	var changes int
	var points []int
	points = make([]int, s.Length)
	changes = 1
	for i := 0; changes > 0; i++ {
		//max iterator
		changes = 0
		// glog.Infof("%s", c.SliceString())
		groups := make(map[int][]*container.Tensor, k)
		vs := s.Values()
		var nearestCI int
		for i, v := range vs {
			nearestCI = nearestOfKMeans(v, c)
			if groups[nearestCI] == nil {
				groups[nearestCI] = make([]*container.Tensor, 0)
			}
			// glog.Infof("%d \n %v \n %v", nearestCI, c, v)
			groups[nearestCI] = append(groups[nearestCI], v)
			if points[i] != nearestCI {
				points[i] = nearestCI
				changes++
			}
		}

		for i, group := range groups {
			if len(group) == 0 {
				// During the iterations, if any of the cluster centers has no
				// data points associated with it, assign a random data point
				// to it.
				// Also see: http://user.ceng.metu.edu.tr/~tcan/ceng465_f1314/Schedule/KMeansEmpty.html
				var ri int
				for {
					// find a cluster with at least two data points, otherwise
					// we're just emptying one cluster to fill another
					ri = rand.Intn(s.Length)
					if len(groups[points[ri]]) > 1 {
						break
					}
				}
				groups[i] = append(groups[i], s.Get(i))
				points[ri] = i
				changes++
			}
		} //for

		if changes > 0 {
			for i, group := range groups {
				_newTensor, _err := meanTensorOfKMeans(group)
				if _err != nil {
					glog.Errorln(_err)
				}
				if !c.Get(i).Equal(_newTensor) {
					c.Put(i, _newTensor)
				}
			}
		} //if
		if i > iter || changes == 0 {
			glog.Infof("iter: %d \t %d", i, changes)
			out := make(map[int]*container.Tensor, len(groups))
			for k, v := range groups {
				out[k] = container.NewTensor(v)
			}
			return out, c
		}
	} //for
	return nil, nil
}

//nearestOfKMeans
func nearestOfKMeans(v, c *container.Tensor) int {

	cs := c.Values()
	var curDist, minDist float64
	var minIndex int
	var err error
	minDist = 2 << 32
	for i, c0 := range cs {
		curDist, err = distance.Euclidean(v, c0)
		if err != nil {
			glog.Errorln(err)
		}
		// glog.Infof("%v ## %v \n %f \t %f", v, c0, curDist, minDist)
		if curDist < minDist {
			minDist = curDist
			minIndex = i
		}
	}
	return minIndex
}

//meanTensor  takes an array of Nodes and returns a node which represents the average
// value for the provided nodes. This is used to center the centroids within their cluster.
func meanTensorOfKMeans(group []*container.Tensor) (*container.Tensor, error) {
	l := len(group)
	if l == 0 {
		return nil, errors.New("there is no mean for an empty set of points")
	}

	c := make([]float64, group[0].Length)
	for _, t := range group {
		vs := t.Values()
		for i, v := range vs {
			// glog.Infoln(v.VType)
			vv, ok := v.Value.(float64)
			if !ok {
				return nil, container.ErrTypeAssertion
			}
			c[i] += vv
		}
	}
	for i, v := range c {
		c[i] = v / float64(l)
	}

	if len(c) == 1 {
		//1-d
		return container.NewTensor(c[0]), nil
	}
	//2-d
	return container.NewTensor(c), nil

}
