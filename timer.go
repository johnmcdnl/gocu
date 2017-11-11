package gocu

import "time"

type Timer struct {
	Begin    time.Time
	End      time.Time
	Duration time.Duration
}

func (t *Timer)Start(){
	t.Begin = time.Now()
}

func (t *Timer)Stop(){
	t.End = time.Now()
	t.Duration = t.End.Sub(t.Begin)
}