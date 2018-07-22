package common

import (
	"github.com/golang/glog"
)

//Attributes the struct
type Attributes map[string]interface{}

// //Attribute
// type Attribute struct {
// 	Key   string
// 	Value interface{}
// }

//NewAttributes create the Attributes
func NewAttributes() *Attributes {
	var attrs Attributes
	m := make(map[string]interface{}, 0)
	attrs = Attributes(m)
	return &attrs
}

//Put add one attribute v with the key k
func (attrs Attributes) Put(k string, v interface{}) bool {
	attrs[k] = v
	return true
}

//Remove remove one attribute with the key k
func (attrs Attributes) Remove(k string) (interface{}, bool) {
	if _, ok := attrs[k]; ok {
		delete(attrs, k)
	}
	return nil, true
}

//Get get one attribute value with the key k
func (attrs Attributes) Get(k string) interface{} {
	return attrs[k]
}

//GetString get one attribute value with the key k
func (attrs Attributes) GetString(k string) string {
	v, ok := attrs[k]
	if !ok {
		glog.Fatalln("k not exist")
	}
	v1, ok := v.(string)
	if !ok {
		glog.Fatalln("type assertion error")
	}
	return v1
}

//Contains return true if the attrs conatians the key k
func (attrs Attributes) Contains(k string) bool {
	_, ok := attrs[k]
	return ok
}
