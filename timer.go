package gojs

import (
	"time"

	"github.com/dop251/goja"
)

type _Timeout struct {
	ch chan struct{}
}

func (_Timeout) String() string {
	return `Timeout`
}

type _Interval struct {
	ch chan struct{}
}

func (_Interval) String() string {
	return `Interval`
}
func (r *Runtime) initTimer() {
	r.Set(`setTimeout`, func(call goja.FunctionCall) goja.Value {
		cb, ok := goja.AssertFunction(call.Argument(0))
		duration := time.Millisecond * time.Duration(call.Argument(1).ToInteger())

		ret := _Timeout{
			ch: make(chan struct{}),
		}
		e := r.Go(func(worker Worker) {
			timer := time.NewTimer(duration)
			select {
			case <-r.ctx.Done(): // runtime cancel
				if !timer.Stop() {
					<-timer.C
				}
			case <-ret.ch: // cancel async
				if !timer.Stop() {
					<-timer.C
				}
				worker.Submit(nil)
			case <-timer.C: // call cb
				if ok {
					worker.Submit(func() {
						cb(r.ToValue(ret))
					})
				} else {
					worker.Submit(nil)
				}
			}
		})
		if e != nil {
			panic(e)
		}
		return r.ToValue(ret)
	})
	r.Set(`clearTimeout`, func(call goja.FunctionCall) goja.Value {
		if v, ok := call.Argument(0).Export().(_Timeout); ok {
			close(v.ch)
		}
		return goja.Undefined()
	})
	r.Set(`setInterval`, func(call goja.FunctionCall) goja.Value {
		cb, ok := goja.AssertFunction(call.Argument(0))
		duration := time.Millisecond * time.Duration(call.Argument(1).ToInteger())
		ret := _Interval{
			ch: make(chan struct{}),
		}
		e := r.Go(func(worker Worker) {
			timer := time.NewTicker(duration)
			for {
				select {
				case <-r.ctx.Done(): // runtime cancel
					timer.Stop()
					return
				case <-ret.ch: // cancel async
					timer.Stop()
					worker.Submit(nil)
					return
				case <-timer.C: // call cb
					var f func()
					if ok {
						f = func() {
							cb(r.ToValue(ret))
						}
					}
					if worker.Next(f) != nil {
						return
					}
				}
			}
		})
		if e != nil {
			panic(e)
		}
		return r.ToValue(ret)
	})
	r.Set(`clearInterval`, func(call goja.FunctionCall) goja.Value {
		if v, ok := call.Argument(0).Export().(_Interval); ok {
			close(v.ch)
		}
		return goja.Undefined()
	})
}
