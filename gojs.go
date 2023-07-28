package gojs

import (
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"

	"github.com/dop251/goja"
)

type Number[T uint8 | uint16 | uint32 | uint64 | int8 | int16 | int32 | int64] struct {
	v       T
	convert func(value any) (T, error)
}

func (v Number[T]) String() string {
	return fmt.Sprint(T(v.v))
}
func (v Number[T]) ToNumber() T {
	return v.v
}
func (v Number[T]) ToUint64() Number[uint64] {
	return MakeUint64(uint64(v.v))
}
func (v Number[T]) ToUint32() Number[uint32] {
	return MakeUint32(uint32(v.v))
}
func (v Number[T]) ToUint16() Number[uint16] {
	return MakeUint16(uint16(v.v))
}
func (v Number[T]) ToUint8() Number[uint8] {
	return MakeUint8(uint8(v.v))
}
func (v Number[T]) ToInt64() Number[int64] {
	return MakeInt64(int64(v.v))
}
func (v Number[T]) ToInt32() Number[int32] {
	return MakeInt32(int32(v.v))
}
func (v Number[T]) ToInt16() Number[int16] {
	return MakeInt16(int16(v.v))
}
func (v Number[T]) ToInt8() Number[int8] {
	return MakeInt8(int8(v.v))
}
func (v Number[T]) Add(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		}
		ret.v += val
	}
	return
}
func (v Number[T]) Sub(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		}
		ret.v -= val
	}
	return
}
func (v Number[T]) Mul(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		}
		ret.v *= val
	}
	return
}
func (v Number[T]) Div(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		} else if val == 0 {
			e = fmt.Errorf(`%v / 0`, ret.v)
			return
		}
		ret.v /= val
	}
	return
}
func (v Number[T]) Mod(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		} else if val == 0 {
			e = fmt.Errorf(`%v %% 0`, ret.v)
			return
		}
		ret.v %= val
	}

	return
}
func (v Number[T]) DivMod(value any) (t0, t1 Number[T], e error) {
	val, e := v.convert(value)
	if e != nil {
		return
	} else if val == 0 {
		e = fmt.Errorf(`%v /~ 0`, v.v)
		return
	}
	t0.v = v.v / val
	t1.v = v.v % val
	return
}
func (v Number[T]) Cmp(value any) (int, error) {
	val, e := v.convert(value)
	if e != nil {
		return 0, e
	}
	if v.v == val {
		return 0, nil
	} else if v.v < val {
		return -1, nil
	}
	return 1, nil
}
func (v Number[T]) Neg() (ret Number[T]) {
	ret.v = -v.v
	return
}
func (v Number[T]) Not() (ret Number[T]) {
	ret.v = ^v.v
	return
}
func (v Number[T]) Or(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		}
		ret.v |= val
	}
	return
}
func (v Number[T]) Xor(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		}
		ret.v ^= val
	}
	return
}
func (v Number[T]) And(vals ...any) (ret Number[T], e error) {
	ret.v = v.v
	var val T
	for _, value := range vals {
		val, e = v.convert(value)
		if e != nil {
			return
		}
		ret.v &= val
	}
	return
}

func ConvertUint[to uint8 | uint16 | uint32 | uint64](value any) (ret to, e error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.String:
		var v64 uint64
		v64, e = strconv.ParseUint(reflect.ValueOf(value).String(), 10, 64)
		if e != nil {
			return
		}
		ret = to(v64)
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		ret = to(reflect.ValueOf(value).Int())
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		ret = to(reflect.ValueOf(value).Uint())
	case reflect.Float32, reflect.Float64:
		ret = to(reflect.ValueOf(value).Float())
	case reflect.Struct:
		switch v := value.(type) {
		case Number[uint64]:
			ret = to(v.v)
		case Number[uint32]:
			ret = to(v.v)
		case Number[uint16]:
			ret = to(v.v)
		case Number[uint8]:
			ret = to(v.v)
		case Number[int64]:
			ret = to(v.v)
		case Number[int32]:
			ret = to(v.v)
		case Number[int16]:
			ret = to(v.v)
		case Number[int8]:
			ret = to(v.v)
		default:
			e = errors.New(`ConvertUint(` + t.String() + `) invalid`)
			return
		}
	default:
		e = errors.New(`ConvertUint(` + t.String() + `) invalid`)
		return
	}
	return
}
func ConvertInt[to int8 | int16 | int32 | int64](value any) (ret to, e error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.String:
		var v64 int64
		v64, e = strconv.ParseInt(reflect.ValueOf(value).String(), 10, 64)
		if e != nil {
			return
		}
		ret = to(v64)
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		ret = to(reflect.ValueOf(value).Int())
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		ret = to(reflect.ValueOf(value).Uint())
	case reflect.Float32, reflect.Float64:
		ret = to(reflect.ValueOf(value).Float())
	case reflect.Struct:
		switch v := value.(type) {
		case Number[uint64]:
			ret = to(v.v)
		case Number[uint32]:
			ret = to(v.v)
		case Number[uint16]:
			ret = to(v.v)
		case Number[uint8]:
			ret = to(v.v)
		case Number[int64]:
			ret = to(v.v)
		case Number[int32]:
			ret = to(v.v)
		case Number[int16]:
			ret = to(v.v)
		case Number[int8]:
			ret = to(v.v)
		default:
			e = errors.New(`ConvertInt(` + t.String() + `) invalid`)
			return
		}
	default:
		e = errors.New(`ConvertInt(` + t.String() + `) invalid`)
		return
	}
	return
}
func ConvertFloat[to float32 | float64](value any) (ret to, e error) {
	t := reflect.TypeOf(value)
	switch t.Kind() {
	case reflect.String:
		var v64 int64
		v64, e = strconv.ParseInt(reflect.ValueOf(value).String(), 10, 64)
		if e != nil {
			return
		}
		ret = to(v64)
	case reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8:
		ret = to(reflect.ValueOf(value).Int())
	case reflect.Uint64, reflect.Uint32, reflect.Uint16, reflect.Uint8:
		ret = to(reflect.ValueOf(value).Uint())
	case reflect.Float32, reflect.Float64:
		ret = to(reflect.ValueOf(value).Float())
	case reflect.Struct:
		switch v := value.(type) {
		case Number[uint64]:
			ret = to(v.v)
		case Number[uint32]:
			ret = to(v.v)
		case Number[uint16]:
			ret = to(v.v)
		case Number[uint8]:
			ret = to(v.v)
		case Number[int64]:
			ret = to(v.v)
		case Number[int32]:
			ret = to(v.v)
		case Number[int16]:
			ret = to(v.v)
		case Number[int8]:
			ret = to(v.v)
		default:
			e = errors.New(`ConvertFloat(` + t.String() + `) invalid`)
			return
		}
	default:
		e = errors.New(`ConvertFloat(` + t.String() + `) invalid`)
		return
	}
	return
}

func MakeUint64(v uint64) Number[uint64] {
	return Number[uint64]{
		v:       v,
		convert: ConvertUint[uint64],
	}
}
func MakeUint32(v uint32) Number[uint32] {
	return Number[uint32]{
		v:       v,
		convert: ConvertUint[uint32],
	}
}
func MakeUint16(v uint16) Number[uint16] {
	return Number[uint16]{
		v:       v,
		convert: ConvertUint[uint16],
	}
}
func MakeUint8(v uint8) Number[uint8] {
	return Number[uint8]{
		v:       v,
		convert: ConvertUint[uint8],
	}
}
func MakeInt64(v int64) Number[int64] {
	return Number[int64]{
		v:       v,
		convert: ConvertInt[int64],
	}
}
func MakeInt32(v int32) Number[int32] {
	return Number[int32]{
		v:       v,
		convert: ConvertInt[int32],
	}
}
func MakeInt16(v int16) Number[int16] {
	return Number[int16]{
		v:       v,
		convert: ConvertInt[int16],
	}
}
func MakeInt8(v int8) Number[int8] {
	return Number[int8]{
		v:       v,
		convert: ConvertInt[int8],
	}
}

func (r *Runtime) initGOJS() {
	r.opts.registry.RegisterNativeModule(`gojs`, func(r *goja.Runtime, m *goja.Object) {
		o := m.Get("exports").(*goja.Object)

		o.Set(`MaxUint64`, r.ToValue(MakeUint64(math.MaxUint64)))
		o.Set(`MaxUint32`, r.ToValue(MakeUint32(math.MaxUint32)))
		o.Set(`MaxUint16`, r.ToValue(MakeUint32(math.MaxUint16)))
		o.Set(`MaxUint8`, r.ToValue(MakeUint8(math.MaxUint8)))
		o.Set(`MaxInt64`, r.ToValue(MakeInt64(math.MaxInt64)))
		o.Set(`MaxInt32`, r.ToValue(MakeInt32(math.MaxInt32)))
		o.Set(`MaxInt16`, r.ToValue(MakeInt16(math.MaxInt16)))
		o.Set(`MaxInt8`, r.ToValue(MakeInt8(math.MaxInt8)))
		o.Set(`MinInt64`, r.ToValue(MakeInt64(math.MinInt64)))
		o.Set(`MinInt32`, r.ToValue(MakeInt32(math.MinInt32)))
		o.Set(`MinInt16`, r.ToValue(MakeInt16(math.MinInt16)))
		o.Set(`MinInt8`, r.ToValue(MakeInt8(math.MinInt8)))

		o.Set(`Uint64`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   uint64
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertUint[uint64](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeUint64(v))
		})
		o.Set(`Uint32`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   uint32
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertUint[uint32](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeUint32(v))
		})
		o.Set(`Uint16`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   uint16
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertUint[uint16](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeUint16(v))
		})
		o.Set(`Uint8`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   uint8
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertUint[uint8](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeUint8(v))
		})
		o.Set(`Int64`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   int64
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertInt[int64](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeInt64(v))
		})
		o.Set(`Int32`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   int32
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertInt[int32](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeInt32(v))
		})
		o.Set(`Int16`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   int16
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertInt[int16](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeInt16(v))
		})
		o.Set(`Int8`, func(call goja.FunctionCall) goja.Value {
			var (
				val = call.Argument(0)
				v   int8
			)
			if !goja.IsUndefined(val) && !goja.IsNull(val) {
				var e error
				v, e = ConvertInt[int8](val.Export())
				if e != nil {
					panic(r.ToValue(e))
				}
			}
			return r.ToValue(MakeInt8(v))
		})
	})
}
