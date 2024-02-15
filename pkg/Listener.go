package pkg

import (
	"fmt"
	"syscall/js"
)

/*
Wrapper of addEventListener and more

	//Example:
	myListener := NewEventListener(myOjectDOM)
	myListener.AddClick("configure click",func(event MouseEvent){
		//...
	})

Or

	myObjectDOM.Listener.AddClick("configure click",func(event MouseEvent){
		//...
	})
*/
type EventListener struct {
	ob     *ObjectDOM
	clicks map[string]func(MouseEvent)
}

// Remove a listener
func (listener *EventListener) RemoveClick(name string) {
	delete(listener.clicks, name)
}

// Add click function with name
//
//	//Example:
//	myObjectDOM.Listener.AddClick("add functions of app",func(event MouseEvent){
//	// you code
//	})
func (listener *EventListener) AddClick(name string, function func(MouseEvent)) {
	listener.clicks[name] = function
}

// internal !
func (listener *EventListener) registerClick() {
	fmt.Println("puto")
	var call js.Func
	call = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		event := args[0]
		call2 := func() {
			event.Call("preventDefault")
		}
		l := MouseEvent{
			X:              event.Get("x").Int(),
			Y:              event.Get("y").Int(),
			ClientX:        event.Get("clientX").Int(),
			ClientY:        event.Get("clientY").Int(),
			Object:         listener.ob,
			PreventDefault: call2,
		}
		for _, po := range listener.clicks {

			po(l)
		}

		return nil
	}) //click
	listener.ob.object.Call("addEventListener", "click", call)
}

// Get all click function with names
func (o *EventListener) GetClicks() map[string]func(MouseEvent) {
	return o.clicks
}

/*
Create a *EventListener

	//Example:
	lis := NewEventListener(myObjectDOM)
	// ...
*/
func NewEventListener(object *ObjectDOM) *EventListener {

	p := &EventListener{
		ob:     object,
		clicks: make(map[string]func(MouseEvent)),
	}
	p.registerClick()
	return p
}
