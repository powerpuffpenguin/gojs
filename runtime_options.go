package gojs

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/dop251/goja_nodejs/require"
)

func wrapSource(filename string, b []byte, exports bool) string {
	if exports {
		return fmt.Sprintf(`(function(){var exports={};(function(__filename,__dirname){%s})(%q,%q)})()`,
			b,
			filename, filepath.Dir(filename),
		)
	} else {
		return fmt.Sprintf(`(function(__filename,__dirname){%s})(%q,%q)`,
			b,
			filename, filepath.Dir(filename),
		)
	}
}

var defaultOptions = options{
	registry: require.NewRegistry(
		require.WithLoader(func(filename string) (b []byte, e error) {
			b, e = os.ReadFile(filename)
			if e != nil {
				if os.IsNotExist(e) {
					e = require.ModuleFileDoesNotExistError
				}
				return
			}
			b = StringToBytes(wrapSource(filename, b, false))
			return
		}),
	),
	console: true,
	timer:   true,
}

type options struct {
	registry *require.Registry
	console  bool
	timer    bool
	loop     Loop
	ctx      context.Context
}
type Option interface {
	apply(*options)
}

type funcOption struct {
	f func(*options)
}

func (fdo *funcOption) apply(do *options) {
	fdo.f(do)
}
func newFuncOption(f func(*options)) *funcOption {
	return &funcOption{
		f: f,
	}
}

func WithRegistry(registry *require.Registry) Option {
	return newFuncOption(func(o *options) {
		o.registry = registry
	})
}

func WithConsole(console bool) Option {
	return newFuncOption(func(o *options) {
		o.console = console
	})
}
func WithTimer(timer bool) Option {
	return newFuncOption(func(o *options) {
		o.timer = timer
	})
}
func WithLoop(loop Loop) Option {
	return newFuncOption(func(o *options) {
		o.loop = loop
	})
}

func WithContext(ctx context.Context) Option {
	return newFuncOption(func(o *options) {
		o.ctx = ctx
	})
}
