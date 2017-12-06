### RingCron

- 实现秒级的环形定时任务触发
- 可以设置任何时长的任务，任何时间执行 

#### 1. install 
```swagger codegen
go get github.com/newgoo/ringcron
```

#### 2. Getting Started
* 创建一个新的定时环任务`New()`
```swagger codegen
r := New(4, time.Second)
```
* 写入一个任务 `InsertTask`
```
r.InsertTask("name", 1, 0, func() {})
```
* 删除一个任务 `RemoveTask`
```
r.RemoveTask("name2")
```

