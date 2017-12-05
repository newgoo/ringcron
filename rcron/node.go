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
	taskList  []*taskSingle
}

func (n *Node) insert(key string, times int, intervalTime time.Duration, f func()) {
	if !n.hasTask {
		n.hasTask = true
	}
	n.lock = new(sync.RWMutex)
	for i, task := range n.taskList {
		if task != nil && task.key == key {
			n.lock.Lock()
			n.taskList[i] = &taskSingle{key: key, task: f, state: true, times: times, intervalTime: intervalTime}
			n.lock.Unlock()
			return
		}
	}

	n.lock.Lock()
	n.taskList = append(n.taskList, &taskSingle{key: key, task: f, state: true, times: times, intervalTime: intervalTime})
	n.lock.Unlock()

}

func (n *Node) execTask() {
	taskLs := make([]*taskSingle, 0)
	for _, task := range n.taskList {
		if task != nil && n.hasTask {
			if task.intervalTime/n.ringTime > 0 {
				task.intervalTime -= n.ringTime
				continue
			}
			go task.task()
			task.state = false
			continue
		}
		taskLs = append(taskLs, task)
	}
	n.taskList = taskLs
	if len(n.taskList) == 0 {
		n.hasTask = false
	}
}

//func (n *Node) execTask() {
//	i := -1
//	for _, task := range n.taskList {
//		i++
//		if task != nil && task.state == true {
//			if task.intervalTime/n.ringTime > 0 {
//				task.intervalTime -= n.ringTime
//				continue
//			}
//			go task.task()
//			task.times--
//			if task.times == 0 {
//				//	log.Info("=========")
//				task.state = false
//				n.lock.Lock()
//				n.taskList = append(n.taskList[:i], n.taskList[i+1:]...)
//				if len(n.taskList) == 0 {
//					n.hasTask = false
//				}
//				n.lock.Unlock()
//				i--
//			}
//
//			//if task.times == 0 {
//			//	n.lock = nil
//			//}
//
//		}
//	}
//}

//func (n *Node) removeTask(key) {
//
//}
