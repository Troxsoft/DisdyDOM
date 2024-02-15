package pkg

import (
	"syscall/js"
)

type IObjectDOM interface {
	Value() string
	SetValue(string) string

	InnerHTML() string
	SetInnerHTML(string)

	GetHTMLType() string

	InnerText() string
	SetInnerText(string)

	AddClass(string)
	Class() []string

	SetID(string)
	ID() string

	RemoveClass(string)

	ClearClass()
	ClearID()

	CallFunction(string, ...interface{})
	GetProperty(string) js.Value
	GetPropertyAsString(string)
	SetProperty(string, ...interface{})
}
type ObjectDOM struct {
	object   js.Value
	Listener *EventListener
}

func (p *ObjectDOM) NewStyle() CSStyle {
	return CSStyle{
		ele:   *p,
		style: p.object.Get("style"),
	}
}
func (object *ObjectDOM) GetHTMLType() string {
	p := object.object.Get("tagName").String()
	return p
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

/*
Value() string
SetValue(string) string

InnerHTML() string
SetInnerHTML(string)

# GetHTMLType() string

InnerText() string
SetInnerText(string)

AddClass(string)
Class() []string

SetID(string)
ID() string

RemoveClass(string)

ClearClass()
ClearID()
*/
func (oj *ObjectDOM) SetID(id string) {

	oj.object.Set("id", id)
}
func (oj *ObjectDOM) ID() string {
	k := oj.object.Get("id").String()
	return k
}
func (ob *ObjectDOM) ClearID() {
	ob.SetID("")
}
func (oj *ObjectDOM) ClearClass() {
	k := oj.Class()

	for _, i := range k {
		oj.RemoveClass(i)
	}
}
func (oj *ObjectDOM) RemoveClass(class string) {
	oj.object.Get("classList").Call("remove", class)
}
func (oj *ObjectDOM) AddClass(class string) {
	oj.object.Get("classList").Call("add", class)
}
func (oj *ObjectDOM) Class() []string {
	classes := oj.object.Get("classList")
	class := []string{}
	// Iterar sobre las clases y obtener sus nombres
	for i := 0; i < classes.Length(); i++ {
		class2 := classes.Index(i).String()
		class = append(class, class2)
	}

	return class

}

/*
CallFunction(string,...interface{})

	GetProperty(string)js.Value
	GetPropertyAsString(string)
	SetProperty(string,...interface{})
*/
func (o *ObjectDOM) CallFunction(function string, params ...interface{}) {
	o.object.Call(function, params)
}
func (o *ObjectDOM) SetProperty(property string, params ...interface{}) {
	o.object.Set(property, params)
}
func (o *ObjectDOM) GetAsStringProperty(property string) string {
	return o.object.Get(property).String()
}
func (o *ObjectDOM) GetProperty(property string) js.Value {
	return o.object.Get(property)
}
