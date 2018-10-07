package common

import (
	"encoding/binary"
	"net"
	"os"
)

//ObjectID (80) define the objectid struct
type ObjectID struct {
	Machine uint32
	Process uint16
	Local   uint32
}

//NewObjectId
func NewObjectID(i uint32) *ObjectID {
	oi := &ObjectID{}
	oi.Machine = MachineIP()
	oi.Process = ProcessID()
	oi.Local = i
	return oi
}

//MachineIP
func MachineIP() uint32 {
	ifaces, err := net.Interfaces()
	if err != nil {
		return uint32(0)
	}
	var ui uint32
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		ui = localIP(iface)
		if ui == 0 {
			continue
		} else {
			return ui
		}
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		ui = localMac(iface)
		if ui == 0 {
			continue
		} else {
			return ui
		}
	}
	return uint32(0)
}

//Process
func ProcessID() uint16 {
	return uint16(os.Getpid())
}

//Equal return true if two ObjectId are same
func (oi *ObjectID) Equal(other *ObjectID) bool {
	if oi.Machine != other.Machine {
		return false
	}
	if oi.Process != other.Process {
		return false
	}
	if oi.Local != other.Local {
		return false
	}
	return true
}

//localIP
func localIP(iface net.Interface) uint32 {
	addrs, err := iface.Addrs()
	if err != nil {
		return uint32(0)
	}
	for _, addr := range addrs {
		var ip net.IP
		switch v := addr.(type) {
		case *net.IPNet:
			ip = v.IP
		case *net.IPAddr:
			ip = v.IP
		}
		// glog.Infof("%v", ip)
		if ip == nil || ip.IsLoopback() {
			continue
		}
		ip = ip.To4()
		if ip == nil {
			continue // not an ipv4 address
		}
		return binary.BigEndian.Uint32([]byte(ip))
	}
	return uint32(0)
}

//localMac
func localMac(iface net.Interface) uint32 {
	if iface.HardwareAddr[0]&2 == 2 {
		return uint32(0)
	}

	_mac := []byte(iface.HardwareAddr)
	// glog.Infof("%v", _mac[2:])
	return binary.BigEndian.Uint32(_mac[2:])
}
