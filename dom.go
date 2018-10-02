package dom

import (
	"syscall/js"

	"github.com/gernest/vected"
)

// Dom implements vected.Value interface which is based on read Dom.
type Dom struct {
	value js.Value
}

func New(v js.Value) *Dom {
	return &Dom{value: v}
}

func (d Dom) Bool() bool {
	return d.value.Bool()
}

func (d Dom) Call(m string, args ...interface{}) vected.Value {
	return New(d.value.Call(m, toArgs(args...)...))
}

func (d Dom) Float() float64 {
	return d.value.Float()
}

func (d Dom) Get(key string) vected.Value {
	defer func() {
		if e := recover(); e != nil {
			Log(d.value)
		}
	}()
	return New(d.value.Get(key))
}

func (d Dom) Index(i int) vected.Value {
	return New(d.value.Index(i))
}

func (d Dom) Int() int {
	return d.value.Int()
}

func (d Dom) Invoke(args ...interface{}) vected.Value {
	return New(d.value.Invoke(toArgs(args...)...))
}

func toArgs(args ...interface{}) (a []interface{}) {
	for _, v := range args {
		if o, ok := v.(*Dom); ok {
			a = append(a, o.value)
			continue
		}
		a = append(a, js.ValueOf(v))
	}
	return
}

func (d Dom) Set(p string, x interface{}) {
	d.value.Set(p, x)
}

func (d Dom) String() string {
	return d.value.String()
}

func (d Dom) Type() vected.Type {
	return vected.Type(d.value.Type())
}

var console = js.Global().Get("console")

func Log(v ...interface{}) {
	console.Call("log", v...)
}
