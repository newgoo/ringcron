package main

import (
	"sync"

	"github.com/lunny/log"
)

func main() {
	mp := new(sync.Map)
	mp.Store("name1", "name")
	mp.Store("name2", "name")
	mp.Store("name3", "name")
	mp.Store("name4", "name")
	mp.Store("name5", "name")

	mp.Range(func(key, value interface{}) bool {
		log.Info(key, value)
		mp.Delete(key)
		mp.Range(func(key, value interface{}) bool {
			log.Info(key, value)
			return true
		})
		//log.Info(key, value)
		return true
	})
}

//var f = func(key, value interface{}) bool {

//}
