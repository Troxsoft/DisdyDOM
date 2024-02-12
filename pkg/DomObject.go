package pkg

import "syscall/js"

type IObjectDOM interface {
	Value() string
	SetValue(string) string

	InnerHTML() string
	SetInnerHTML(string)

	InnerText() string
	SetInnerText(string)
}

type ObjectDOM struct {
	object js.Value
}

func (object *ObjectDOM) SetInnerText(value string) {
	object.object.Set("innerText", value)
}
func (object *ObjectDOM) InnerText() string {
	p := object.object.Get("innerText").String()
	return p
}

func (object *ObjectDOM) SetInnerHTML(value string) {
	object.object.Set("innerHTML", value)
}
func (object *ObjectDOM) InnerHTML() string {
	p := object.object.Get("innerHTML").String()
	return p
}

func (object *ObjectDOM) SetValue(value string) {
	object.object.Set("innerText", value)
}
func (object *ObjectDOM) Value() string {
	p := object.object.Get("innerText").String()
	return p
}
