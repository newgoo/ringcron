package rcron

import (
	"container/ring"
	"fmt"
	"sync"
	"time"

	"github.com/lunny/log"
)

type RCron struct {
	ring      *ring.Ring
	isAllow   bool
	currentId int
	nodes     map[int]*Node
	duration  time.Duration
	keys      *sync.Map
}

func New(len int, duration time.Duration) (rCron *RCron) {
	rCron = new(RCron)
	rCron.keys = new(sync.Map)
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

func (r *RCron) InsertTask(key string, times int, intervalTime time.Duration, f interface{}, params ...interface{}) {
	if times != 1 {
		return
	}
	crId := r.currentId
	nodeId := (crId + 1 + int((intervalTime%(r.duration*time.Duration(r.ring.Len())))/r.duration)) % r.ring.Len()
	r.keys.Store(key, nodeId)
	node := r.nodes[nodeId]
	node.insert(key, times, intervalTime, f, params...)
}

func (r *RCron) RemoveTask(key string) (err error) {
	value, ok := r.keys.Load(key)
	if !ok {
		return fmt.Errorf("Task not exist! ")
	}
	nodeId, ok := value.(int)
	if !ok {
		return fmt.Errorf("Error ")
	}
	r.nodes[nodeId].delTask(key)
	r.keys.Delete(key)
	return
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
