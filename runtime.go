package gojs

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/dop251/goja"
	"github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/require"
)

type Runtime struct {
	*goja.Runtime
	opts   options
	ctx    context.Context
	cancel context.CancelFunc
}

func New(opt ...Option) *Runtime {
	ret := &Runtime{
		Runtime: goja.New(),
		opts:    defaultOptions,
	}
	for _, o := range opt {
		o.apply(&ret.opts)
	}
	if ret.opts.ctx == nil {
		ret.ctx, ret.cancel = context.WithCancel(context.Background())
	} else {
		ret.ctx, ret.cancel = context.WithCancel(ret.opts.ctx)
	}
	ret.opts.registry.Enable(ret.Runtime)
	if ret.opts.loop == nil {
		ret.opts.loop = NewSimpleLoop()
	}
	if ret.opts.console {
		ret.initConsole()
	}
	if ret.opts.timer {
		ret.initTimer()
	}
	ret.initGOJS()
	return ret
}
func (r *Runtime) Cancel() {
	r.cancel()
}
func (r *Runtime) GetRegistry() *require.Registry {
	return r.opts.registry
}
func (r *Runtime) GetLoop() Loop {
	return r.opts.loop
}
func (r *Runtime) RunLoop() error {
	return r.opts.loop.Run(r.ctx)
}
func (r *Runtime) Go(job func(worker Worker)) error {
	return r.opts.loop.Go(r.ctx, job)
}
func (r *Runtime) RunScript(name string, src interface{}) (v goja.Value, e error) {
	if filepath.IsAbs(name) {
		name = filepath.Clean(name)
	} else {
		name, e = filepath.Abs(name)
		if e != nil {
			return
		}
	}
	b, e := parser.ReadSource(name, src)
	if e != nil {
		return
	}
	return r.Runtime.RunScript(name, wrapSource(name, b, true))
}
func (r *Runtime) initConsole() {
	obj := r.NewObject()
	r.Set(`console`, obj)
	obj.Set(`trace`, r.nativeLog)
	obj.Set(`log`, r.nativeLog)
	obj.Set(`info`, r.nativeLog)
	obj.Set(`debug`, r.nativeLog)
	obj.Set(`error`, r.nativeLog)
}
func (r *Runtime) nativeLog(call goja.FunctionCall) goja.Value {
	args := make([]interface{}, len(call.Arguments))
	for i, arg := range call.Arguments {
		if v, ok := arg.Export().(interface{ String() string }); ok {
			args[i] = v.String()
		} else {
			args[i] = arg
		}
	}
	fmt.Println(args...)
	return goja.Undefined()
}
