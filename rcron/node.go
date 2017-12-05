package rcron

import (
	"sync"
	"time"
)

type Node struct {
	ringTime  time.Duration
	lock      *sync.RWMutex
	ringValue interface{}
	hasTask   bool
	taskList  []*task
}

func (n *Node) insert(key string, times int, intervalTime time.Duration, f func()) {
	if !n.hasTask {
		n.hasTask = true
	}
	n.taskList = append(n.taskList, &task{key: key, task: f, state: true, times: times, intervalTime: intervalTime})
	n.lock = new(sync.RWMutex)
}

func (n *Node) execTask() {
	for i, task := range n.taskList {
		if task.times > 0 && task.state == true {
			if task.intervalTime/n.ringTime > 0 {
				task.intervalTime -= n.ringTime
				continue
			}
			task.task()
			n.lock.Lock()
			task.times--
			if task.times == 0 {
				task.state = false
				n.taskList = append(n.taskList[:i], n.taskList[i+1:]...)
				if len(n.taskList) == 0 {
					n.hasTask = false
				}
			}
			n.lock.Unlock()
			if task.times == 0 {
				n.lock = nil
			}

		}
	}
}
