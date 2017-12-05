package rcron

import "time"

type task struct {
	state        bool
	key          string
	times        int
	intervalTime time.Duration
	task         func()
}
