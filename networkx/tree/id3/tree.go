package id3

import (
	"fmt"
	"math"
	"strings"

	"github.com/golang/glog"
	"mervin.me/algset/container"
)

/*
以信息论为基础，以信息熵和信息增益为衡量标准，从而实现对数据的归纳分类。

1. 选择信息增益最大的属性作为当前的特征对数据集分类

    ### 信息熵(Information entropy)

    $$
    Entropy(S) = - \sum_{i=1}^{m}{p(u_i)log_2p(u_i)}
    $$

    其中，$p(u_i) = \frac{|u_i|}{|S|}$,为类别$u_i$在样本空间$S$中出现的概率

    ### 信息增益(Information gain)

    $$
    G(S, A)  = Entropy(S) - \sum_{V\in Value(A)} {\frac{|S_V|}{|S|}Entropy(S_V)}
    $$

    其中，$A$表示样本属性，$Value(A)$是属性A所有的取值集合，$S_V$是$S$中$A$的值为$V$的样例集合


2. 划分的结束：
    划分出来的类属于同一个类
    没有属性可供再分

Ref:
1. https://blog.csdn.net/jerry81333/article/details/53125197/
*/

const (
	//AttributeLabelKey define the key of the AttributeIndexKey attribute
	AttributeLabelKey string = "Node_Label"
	//EntropyGainKey define the EG key of the EntropyGainKey attribute
	EntropyGainKey string = "Node_EG"
	//AttributeValueKey define the key of the AttributeValueKey attribute
	AttributeValueKey string = "Node_Value"
	//ClassificationLableKey define the key of the  ClassificationLableKey attribute
	ClassificationLableKey string = "Class_Label"
)

//DecisionTree the struct of the ID3 DecisionTree
type DecisionTree struct {
	Root *Node

	data  *container.Tensor
	label []string
}

//NewDecisionTree create a new ID3 DecisionTree
func NewDecisionTree(data *container.Tensor, label []string) *DecisionTree {
	dt := new(DecisionTree)
	dt.data = data
	dt.label = label
	dt.creatNode(nil, data, label)
	return dt
}

//creatNode create a new node for the best attribute
func (dt *DecisionTree) creatNode(parent *Node, data *container.Tensor, label []string) *Node {
	var r *Node
	tmp := label[0]
	leafFlag := true
	for _, v := range label {
		if !strings.EqualFold(tmp, v) {
			leafFlag = false
			break
		}
	}
	if leafFlag {
		r = NewNode(fmt.Sprintf("%d", -1))
		// glog.Infof("leaf: %s", label[0])
		r.Attributes.Put(ClassificationLableKey, label[0])
		r.Attributes.Put(AttributeLabelKey, "-1")
		if parent == nil {
			dt.Root = r
		} else {
			r.parent = parent
		}
		return r
	}
	// glog.Infof("%v \n %v", data.SliceString(), label)
	index, eg, freq := chooseBestAttribute(data, label)
	// glog.Infof("best attribute: %s", data.Get(index).Label)
	// glog.Infof("%v", freq)
	r = NewNode(fmt.Sprintf("%d", index))
	r.Attributes.Put(EntropyGainKey, eg)

	if parent == nil {
		dt.Root = r
	} else {
		r.parent = parent
	}

	r.Attributes.Put(AttributeLabelKey, data.Get(index).Label)
	//creat edges
	dt.creatEdge(r, index, data, label, freq)
	// }

	return r

}

//creatEdge create all the edges from the parent to others
func (dt *DecisionTree) creatEdge(source *Node, index int, data *container.Tensor, label []string, freq map[float64][]int) {
	rIndex := make([]int, data.Length-1)
	j := 0
	for i := 0; i < data.Length; i++ {
		if i == index {
			continue
		}
		rIndex[j] = i
		j++
	}
	var subData *container.Tensor
	var subLabel []string
	var target *Node
	var e *Edge
	for k, vs := range freq {
		subLabel = make([]string, len(vs))
		for i, v1 := range vs {
			subLabel[i] = label[v1]
		} //for
		subData = data.Filter(rIndex, vs)
		target = dt.creatNode(source, subData, subLabel)
		e = source.AddEdge(target)
		e.Attributes.Put(AttributeValueKey, float64(k))
	}
}

//BFS the BFS of the tree
func (dt *DecisionTree) BFS() {

	q := container.NewQueue()
	q.Push(dt.Root)
	for !q.Empty() {
		item, _ := q.Pop()
		n, _ := item.(*Node)
		if len(n.children) > 0 {
			for _, c := range n.children {
				glog.Infof("%s - %s", c.Target.Label, c.Target.Attributes.Get(AttributeLabelKey))
				q.Push(c.Target)
			}
			glog.Infoln("################")
		}
	}
}

//chooseBestAttribute return the attribute index with the maximum information gain
func chooseBestAttribute(data *container.Tensor, label []string) (int, float64, map[float64][]int) {
	//data: 2-d
	//the information entorpy for the data
	e := entropy(label)
	dataLen := len(label)
	vs := data.Values()

	var vf1, vf2, ve float64
	var err error
	var freqOfOneDim1 map[float64][]string
	var freqOfOneDim2 map[float64][]int
	var freqOfMaxEG map[float64][]int

	var maxEGValue float64
	var maxEGIndex int
	maxEGValue = 0
	maxEGIndex = -1
	for i, v := range vs {
		// the values(or time series) of the attribuate xx
		freqOfOneDim1 = make(map[float64][]string, 0)
		freqOfOneDim2 = make(map[float64][]int, 0)
		//6 decimal places
		vs1 := v.Values()
		vf1 = 0
		ve = 0
		for i1, v1 := range vs1 {

			//get the value of v1
			vf1, err = v1.Float64()
			if err != nil {
				glog.Fatalln(err)
			}
			vf2 = float64(int(vf1*1000000)) / 1000000
			// glog.Infof("%f-%f", vf1, vf2)
			//classify the v1 value
			if _, ok := freqOfOneDim1[vf1]; ok {
				freqOfOneDim1[vf2] = append(freqOfOneDim1[vf2], label[i1])
				freqOfOneDim2[vf2] = append(freqOfOneDim2[vf2], i1)
			} else {
				freqOfOneDim1[vf2] = []string{label[i1]}
				freqOfOneDim2[vf2] = []int{i1}
			}
		} //for
		//subplace entropy
		for _, v2 := range freqOfOneDim1 {
			// glog.Infof("%f %d %d ", k, len(v2), dataLen)
			// glog.Infof("%v", v2)
			ve += float64(len(v2)) / float64(dataLen) * entropy(v2)
			// glog.Infof("%f %f %f", entropy(v2), float64(len(v2))/float64(dataLen)*entropy(v2), ve)
		}
		// glog.Infof("%s: %f \t %f \t %f", v.Label, e, ve, e-ve)
		if maxEGValue < (e - ve) {
			maxEGValue = e - ve
			maxEGIndex = i
			freqOfMaxEG = freqOfOneDim2
		} //if
	} //if

	//get the maximum information gain
	return maxEGIndex, maxEGValue, freqOfMaxEG
}

//entropy calcaulte the information entropy
func entropy(label []string) float64 {

	l := len(label)
	freq := make(map[string]int, 0)
	for _, v := range label {
		if _, ok := freq[v]; ok {
			freq[v]++
		} else {
			freq[v] = 1
		}
	}
	var e float64
	for _, v := range freq {
		e -= (float64(v) / float64(l)) * math.Log2(float64(v)/float64(l))
	}
	return e
}
