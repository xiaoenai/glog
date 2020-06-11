package glog

import (
	"os"
	"testing"
)

func TestExampleLog(t *testing.T) {
	Printf("print")
	Tracef("trace")
	Debugf("debug")
	Infof("info")
	Errorf("error")
	Fatalf("fatal")

	DefaultLogger().Level = INFO

	Tracef("trace")
	Debugf("debug")
	Infof("info")
	Errorf("error")
	Fatalf("fatal")

	DefaultLogger().Colorful = false

	Tracef("trace")
	Debugf("debug")
	Infof("info")
	Errorf("error")
	Fatalf("fatal")

	NewNamedLogger("named", os.Stdout)
	nl := GetLogger("named")
	nl.Tracef("trace, %d", 1)
	nl.Debugf("debug, %d", 2)
	nl.Infof("info, %d", 3)
	nl.Errorf("error, %d", 4)
	nl.Fatalf("fatal, %d", 5)

	sl := nl.SetPrefix("STAT")
	sl.Tracef("trace, %d", 1)
	sl.Debugf("debug, %d", 2)
	sl.Infof("info, %d", 3)
	sl.Errorf("error, %d", 4)
	sl.Fatalf("fatal, %d", 5)
}
