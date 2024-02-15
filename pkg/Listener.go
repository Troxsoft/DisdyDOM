package pkg

import (
	"fmt"
	"syscall/js"
)

type EventListener struct {
	ob     *ObjectDOM
	clicks map[string]func(MouseEvent)
}

func (listener *EventListener) RemoveClick(name string) {
	delete(listener.clicks, name)
}
func (listener *EventListener) AddClick(name string, function func(MouseEvent)) {
	listener.clicks[name] = function
}
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
func (o *EventListener) GetClicks() map[string]func(MouseEvent) {
	return o.clicks
}
func NewEventListener(object *ObjectDOM) *EventListener {

	p := &EventListener{
		ob:     object,
		clicks: make(map[string]func(MouseEvent)),
	}
	p.registerClick()
	return p
}
