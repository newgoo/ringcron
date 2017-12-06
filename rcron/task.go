package rcron

import (
	"reflect"
	"time"
)

type taskSingle struct {
	//state        bool
	key          string
	times        int
	intervalTime time.Duration
	task         interface{}
	params       []reflect.Value
}
