package rcron

import (
	"testing"
	"time"

	"fmt"

	"github.com/lunny/log"
)

func TestNew(t *testing.T) {
	r := New(4, time.Second)
	go r.Exec()
	time.Sleep(time.Second * 2)
	r.InsertTask("name", 1, 0, func() {})
	time.Sleep(time.Second * 2)
	r.InsertTask("name2", 1, 2, func() {
		log.Infof("---------2----------------")
	})
	log.Info(r.RemoveTask("name1"))

	r.InsertTask("name2", 1, 0, func() {
		log.Infof("---------3----------------")
	})
	log.Info(r.RemoveTask("name2"))
	r.InsertTask("name4", 1, 0, func() {
		log.Infof("---------4----------------")
	})
	time.Sleep(time.Minute)
}

func BenchmarkNew(b *testing.B) {
	r := New(10, time.Second)
	b.ResetTimer()
	go r.Exec()
	for i := 0; i < b.N; i++ {
		r.InsertTask(fmt.Sprintf("%d", i), 1, 0, func() {
			log.Info("==============", i)
		})
	}
}
