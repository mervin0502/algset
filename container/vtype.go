package container

import "reflect"

//TensorVType
type VType uint8

const (
	VBool VType = iota
	VInt
	VInt8
	VInt16
	VInt32
	VInt64
	VUint
	VUint8
	VUint16
	VUint32
	VUint64
	VFloat32
	VFloat64
	VString
	VComplex64
	VComplex128

	VSlice
	VMap
	VTensor
	Nil
)

//VTypeOf return the vtype of v
func VTypeOf(v interface{}) VType {
	t := reflect.TypeOf(v)
	var vt VType
	switch t.Kind() {
	case reflect.Bool:
		vt = VBool
		break
	case reflect.Int:
		vt = VInt
		break

	case reflect.Int8:
		vt = VInt8
		break
	case reflect.Int16:
		vt = VInt16
		break

	case reflect.Int32:
		vt = VInt32
		break
	case reflect.Int64:
		vt = VInt64
		break
	case reflect.Uint8:
		vt = VUint8
		break
	case reflect.Uint16:
		vt = VUint16
		break
	case reflect.Uint32:
		vt = VUint32
		break
	case reflect.Uint64:
		vt = VUint64
		break
	case reflect.Float32:
		vt = VFloat32
		break
	case reflect.Float64:
		vt = VFloat64
		break
	case reflect.String:
		vt = VString
		break
	case reflect.Complex64:
		vt = VComplex64
		break
	case reflect.Complex128:
		vt = VComplex128
		break

	default:
		vt = Nil
	}
	return vt
}

//String
func (v VType) String() string {
	var vt string
	switch v {
	case VBool:
		vt = "VBool"
		break
	case VInt:
		vt = "VInt"
		break

	case VInt8:
		vt = "VInt8"
		break
	case VInt16:
		vt = "VInt16"
		break

	case VInt32:
		vt = "VInt32"
		break
	case VInt64:
		vt = "VInt64"
		break
	case VUint8:
		vt = "VUint8"
		break
	case VUint16:
		vt = "VUint16"
		break
	case VUint32:
		vt = "VUint32"
		break
	case VUint64:
		vt = "VUint64"
		break
	case VFloat32:
		vt = "VFloat32"
		break
	case VFloat64:
		vt = "VFloat64"
		break
	case VString:
		vt = "VString"
		break
	case VComplex64:
		vt = "VComplex64"
		break
	case VComplex128:
		vt = "VComplex128"
		break

	default:
		vt = "Nil"
	}
	return vt
}
