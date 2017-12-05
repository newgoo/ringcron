package rcron

import (
	"time"
)

type taskSingle struct {
	state        bool
	key          string
	times        int
	intervalTime time.Duration
	task         func()
}
