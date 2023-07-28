package gojs

import (
	"context"
	"errors"
	"sync/atomic"
)

type Loop interface {
	// Enable an asynchronous event
	Go(ctx context.Context, job func(worker Worker)) error
	// Run the event loop until all events complete
	Run(ctx context.Context) (e error)
}
type Worker interface {
	// The submission event is completed.
	// If the event system has been destroyed or repeated submission will return an error,
	// you should consider releasing the resource at this time (no other callback function will get the resource and release it)
	Submit(f func()) error
	// Call back data to the event system, but the event has not yet completed
	Next(f func()) error
}

type SimpleLoop struct {
	ch   chan _Next
	jobs int64
}

func NewSimpleLoop() *SimpleLoop {
	return &SimpleLoop{
		ch: make(chan _Next),
	}
}

type _Next struct {
	f    func()
	done bool
}
type _Worker struct {
	submit int32
	ctx    context.Context
	ch     chan<- _Next
}

func (w *_Worker) Submit(f func()) error {
	if w.submit == 0 && atomic.SwapInt32(&w.submit, 1) == 0 { // only submit once
		select {
		case <-w.ctx.Done():
			return w.ctx.Err()
		case w.ch <- _Next{
			f:    f,
			done: true,
		}:
			return nil
		}
	} else {
		return errors.New(`repeated submit`)
	}
}
func (w *_Worker) Next(f func()) error {
	if w.submit == 0 && atomic.LoadInt32(&w.submit) == 0 { // only submit once
		select {
		case <-w.ctx.Done():
			return w.ctx.Err()
		case w.ch <- _Next{
			f:    f,
			done: false,
		}:
			return nil
		}
	} else {
		return errors.New(`repeated submit`)
	}
}
func (s *SimpleLoop) Go(ctx context.Context, job func(worker Worker)) error {
	e := ctx.Err()
	if e != nil {
		return e
	}

	atomic.AddInt64(&s.jobs, 1)

	worker := &_Worker{
		ctx: ctx,
		ch:  s.ch,
	}
	go job(worker)
	return nil
}
func (s *SimpleLoop) Run(ctx context.Context) (e error) {
	var done = ctx.Done()
	jobs := atomic.LoadInt64(&s.jobs)
	for jobs != 0 {
		select {
		case <-done:
			e = ctx.Err()
			return
		case next := <-s.ch:
			if next.f != nil {
				next.f()
			}
			if next.done {
				jobs = atomic.AddInt64(&s.jobs, -1)
			}
		}
	}
	return
}
