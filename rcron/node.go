package rcron

import (
	"sync"
	"time"
)

type Node struct {
	ringTime time.Duration
	//lock      *sync.RWMutex
	ringValue interface{}
	//Count     int
	taskMap *sync.Map
}

func (n *Node) insert(key string, times int, intervalTime time.Duration, f func()) {
	n.taskMap.Store(key, &taskSingle{key: key, task: f, times: times, intervalTime: intervalTime})
}

//func (n *Node) insert(key string, times int, intervalTime time.Duration, f func()) {
//	n.taskMap.Store(key, &taskSingle{key: key, task: f, state: true, times: times, intervalTime: intervalTime})
//}

func (n *Node) execTask() {
	n.taskMap.Range(n.exec)
}

func (n *Node) exec(key, value interface{}) bool {
	var tsk *taskSingle
	var ok bool
	if tsk, ok = value.(*taskSingle); !ok {
		return true
	}

	if tsk != nil {
		if tsk.intervalTime/n.ringTime > 0 {
			tsk.intervalTime -= n.ringTime
			return true
		}
		signal := make(chan interface{})
		go n.runTask(signal, key, tsk.task)
		go n.deleteTask(signal)
		//tsk.state = false
	}
	return true
}

func (n *Node) runTask(signal chan interface{}, key interface{}, f func()) {
	f()
	signal <- key
}

func (n *Node) deleteTask(signal chan interface{}) {
	n.taskMap.Delete(<-signal)
}

//list
//func (n *Node) execTask() {
//	taskLs := make([]*taskSingle, 0)
//	for _, task := range n.taskList {
//		if task != nil && n.hasTask {
//			if task.intervalTime/n.ringTime > 0 {
//				task.intervalTime -= n.ringTime
//				continue
//			}
//			go task.task()
//			task.state = false
//			continue
//		}
//		taskLs = append(taskLs, task)
//	}
//	n.taskList = taskLs
//	if len(n.taskList) == 0 {
//		n.hasTask = false
//	}
//}

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
