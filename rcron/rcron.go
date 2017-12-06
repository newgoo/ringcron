package rcron

import (
	"container/ring"
	"time"

	"sync"

	"github.com/lunny/log"
)

type RCron struct {
	ring      *ring.Ring
	isAllow   bool
	currentId int
	nodes     map[int]*Node
	duration  time.Duration
}

func New(len int, duration time.Duration) (rCron *RCron) {
	rCron = new(RCron)
	rCron.duration = duration
	rCron.ring = ring.New(len)
	rCron.nodes = make(map[int]*Node)

	for i := 0; i < rCron.ring.Len(); i++ {
		node := new(Node)
		node.taskMap = new(sync.Map)
		node.ringTime = duration * time.Duration(len)
		rCron.ring.Value = node
		rCron.nodes[i] = node
		rCron.ring = rCron.ring.Next()
	}
	return
}

func (r *RCron) Exec() {
	for range time.Tick(r.duration) {
		log.Info(r.ring.Value, r.currentId)
		r.ring = r.ring.Next()
		r.currentId = (r.currentId + 1) % r.ring.Len()
		if _, ok := r.ring.Value.(*Node); ok {
			//if nd.hasTask {
			go r.ring.Value.(*Node).execTask()
			//}
		}
	}
}

func (r *RCron) InsertTask(key string, times int, intervalTime time.Duration, f func()) {
	if times != 1 {
		return
	}
	node := r.nodes[(r.currentId+1+int((intervalTime%(r.duration*time.Duration(r.ring.Len())))/r.duration))%r.ring.Len()]
	node.insert(key, times, intervalTime, f)
}

func (r *RCron) RemoveTask() {

}

//func (r *RCron) CloseTask() {
//
//}
//
//func (r *RCron) Move(step int) {
//
//}
//
//func (r *RCron) Link() {
//
//}
