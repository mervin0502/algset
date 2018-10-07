package common

import (
	"flag"
	"testing"
)

func init() {
	flag.Parse()
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", "data/log/")
	flag.Set("v", "2")
}
func TestMachineIP(t *testing.T) {
	t.Logf("%v", MachineIP())
}

func TestProcessID(t *testing.T) {
	t.Logf("%v", ProcessID())
}
