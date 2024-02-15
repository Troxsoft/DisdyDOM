package pkg

import "syscall/js"

// Represents a Global object in Javascript
type WindowDOM struct {
	win js.Value
}

// Show Alert
func (win *WindowDOM) Alert(message string) {
	win.CallFunction("alert", message)
}

// Call global function
func (win *WindowDOM) CallFunction(function string, params ...interface{}) {
	win.win.Call(function, params)
}

// Create a new *WindowDOM
func NewWindowDOM() *WindowDOM {
	return &WindowDOM{
		win: js.Global(),
	}
}
