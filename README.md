# glog
golang log util

## Install

```
go get github.com/xiaoenai/glog
```

## Example

```
package main

import (
	"time"
    "github.com/xiaoenai/glog"
)

func main(){
	glog.DefaultLogger().Level =glog.NewLevel("TRACE")
	glog.Tracef("TRACE-> %d",time.Now().Unix())
	glog.Debugf("DEBUG")
	glog.Infof("INFO")
	glog.Warnf("WARN")
	glog.Errorf("ERROR")
	glog.Fatalf("FATALF")
}
```