package main

import (
	"time"

	"github.com/lunny/log"
	"github.com/newgoo/ringcron/rcron"
)

//var r *rcron.RCron

func main() {
	r := rcron.New(4, time.Second)
	go r.Exec()
	time.Sleep(time.Second * 2)
	r.InsertTask("name", 1, 0, func() {
		log.Infof("--------1-----------------")
	})
	time.Sleep(time.Second * 2)
	r.InsertTask("name2", 1, 2, func() {
		log.Infof("---------2----------------")
	})

	r.InsertTask("name3", 1, 0, func() {
		log.Infof("---------3----------------")
	})
	r.InsertTask("name4", 1, 0, func() {
		log.Infof("---------4----------------")
	})
	time.Sleep(time.Minute)
}
