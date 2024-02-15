package pkg

import "syscall/js"

type WindowDOM struct {
	win js.Value
}

func (win *WindowDOM) Alert(message string) {
	win.CallFunction("alert", message)
}

func (win *WindowDOM) CallFunction(function string, params ...interface{}) {
	win.win.Call(function, params)
}
func NewWindowDOM() *WindowDOM {
	return &WindowDOM{
		win: js.Global(),
	}
}
