package main

import (
	"reflect"

	"github.com/lunny/log"
)

func main() {
	s(f, "张三")
}

var f = func(name string) {
	log.Info("=========", name)
}

func s(f interface{}, name string) {
	//log.Info(reflect.TypeOf(f))
	value := reflect.ValueOf(f)
	in := make([]reflect.Value, 0)
	in = append(in, reflect.ValueOf(name))
	value.Call(in)
}
