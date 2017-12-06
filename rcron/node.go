package rcron

import (
	"fmt"
	"reflect"
	"sync"
	"time"

	"github.com/lunny/log"
)

type Node struct {
	ringTime time.Duration
	//lock      *sync.RWMutex
	//ringValue interface{}
	//Count     int
	taskMap *sync.Map
}

func (n *Node) insert(key string, times int, intervalTime time.Duration, f interface{}, params ...interface{}) (err error) {
	in := make([]reflect.Value, 0)
	if len(params) != reflect.ValueOf(f).Type().NumIn() {
		log.Info("err", len(params), reflect.ValueOf(f).Type().NumIn())
		return fmt.Errorf("error params num ")
	}
	for _, one := range params {
		in = append(in, reflect.ValueOf(one))
	}
	n.taskMap.Store(key, &taskSingle{key: key, task: f, times: times, intervalTime: intervalTime, params: in})
	return
}

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
		go n.runTask(signal, key, tsk.task, tsk.params)
		go n.deleteTask(signal)
		//tsk.state = false
	}
	return true
}

func (n *Node) runTask(signal chan interface{}, key interface{}, f interface{}, in []reflect.Value) {
	value := reflect.ValueOf(f).Call(in)
	fmt.Println(value)
	signal <- key
}

func (n *Node) deleteTask(signal chan interface{}) {
	n.taskMap.Delete(<-signal)
}

func (n *Node) delTask(key interface{}) {
	n.taskMap.Delete(key)
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
