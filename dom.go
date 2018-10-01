package dom

import (
	"syscall/js"

	"github.com/gernest/vected"
)

// Dom implements vected.Value interface which is based on read Dom.
type Dom struct {
	value js.Value
}

func New(v js.Value) Dom {
	return Dom{value: v}
}

func (d Dom) Bool() bool {
	return d.value.Bool()
}

func (d Dom) Call(m string, args ...interface{}) vected.Value {
	return New(d.value.Call(m, args...))
}

func (d Dom) Float() float64 {
	return d.value.Float()
}

func (d Dom) Get(key string) vected.Value {
	return New(d.value.Get(key))
}

func (d Dom) Index(i int) vected.Value {
	return New(d.value.Index(i))
}

func (d Dom) Int() int {
	return d.value.Int()
}

func (d Dom) Invoke(args ...interface{}) vected.Value {
	return New(d.value.Invoke(args...))
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
