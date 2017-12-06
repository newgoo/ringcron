package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f = func(name string) func() {
		return func() {
			fmt.Println("---------", name)
		}
	}
	sm(f("aaa"))
}

func sm(f interface{}) {
	fu, ok := f.(func())
	if ok {
		fu()
	}
	fmt.Println(reflect.TypeOf(f))
}
