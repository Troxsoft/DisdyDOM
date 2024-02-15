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

/*
Represents a DOM. object

	//Example:
	oj := Document.GetByID("you id")
	oj.SetValue("hello developer :)")
*/
type ObjectDOM struct {
	object   js.Value
	Listener *EventListener
}

/*
Create a style bound to the object

	//Example:
	ob.NewStyle().Inject("color","blue")
*/
func (p *ObjectDOM) NewStyle() CSStyle {
	return CSStyle{
		ele:   *p,
		style: p.object.Get("style"),
	}
}

/*
Get the type of the HTML object

	//Example
	jk := Document.Create(H1)
	fmt.Println(jk.GetHTMLType()) // h1
*/
func (object *ObjectDOM) GetHTMLType() string {
	p := object.object.Get("tagName").String()
	return p
}

/*
Setter of innertText (ONLY TEXT)

	// Example
	oj.SetInnerText("hello dev")
*/
func (object *ObjectDOM) SetInnerText(value string) {
	object.object.Set("innerText", value)
}

/*
Getter of innertText (ONLY TEXT)

	// Example
	fmt.Println(oj.InnertText())
*/
func (object *ObjectDOM) InnerText() string {
	p := object.object.Get("innerText").String()
	return p
}

/*
Setter of innertHTML (VALID HTML)

	// Example
	oj.SetInnerHTML("hello dev")
*/
func (object *ObjectDOM) SetInnerHTML(value string) {
	object.object.Set("innerHTML", value)
}

/*
Getter of innertText (VALID HTML)

	// Example
	fmt.Println(oj.InnertHTML())
*/
func (object *ObjectDOM) InnerHTML() string {
	p := object.object.Get("innerHTML").String()
	return p
}

// Alias of SetInnerText
func (object *ObjectDOM) SetValue(value string) {
	object.SetInnerText(value)
}

// Alias of InnerText
func (object *ObjectDOM) Value() string {
	p := object.object.Get("innerText").String()
	return p
}

/*
Setter of id

	//Example:
	information := Document.Create(P)
	information.SetID("info")
*/
func (oj *ObjectDOM) SetID(id string) {

	oj.object.Set("id", id)
}

/*
Getter of id

	//Example:
	information := Document.Create(P)
	information.SetID("info")
	fmt.Println(information.ID()) // info
*/
func (oj *ObjectDOM) ID() string {
	k := oj.object.Get("id").String()
	return k
}

/*Alias of SetID("")*/
func (ob *ObjectDOM) ClearID() {
	ob.SetID("")
}

/*
Remove a all class of ObjectDOM

	//Example:
	oj := Document.Create(H4)
	oj.AddClass("info")
	oj.AddClass("error")
	fmt.Println(oj.Class()) // ["info","error"]
	oj.ClearClass()
	fmt.Println(oj.Class()) // []
*/
func (oj *ObjectDOM) ClearClass() {
	k := oj.Class()

	for _, i := range k {
		oj.RemoveClass(i)
	}
}

/*
Remove a class

	//Example:
	oj := Document.Create(H4)
	oj.AddClass("info")
	oj.AddClass("error")
	fmt.Println(oj.Class()) // ["info","error"]
	oj.RemoveClass("error")
	fmt.Println(oj.Class()) // ["info"]
*/
func (oj *ObjectDOM) RemoveClass(class string) {
	oj.object.Get("classList").Call("remove", class)
}

/*
Add a class to ObjectDOM

	//Example:
	oj := Document.Create(H4)
	oj.AddClass("isGood")
	fmt.Println(oj.Class()) // ["isGood"]
*/
func (oj *ObjectDOM) AddClass(class string) {
	oj.object.Get("classList").Call("add", class)
}

/*
Get a all Class as slice of strings

	//Example:
	oj := Document.Create(H4)
	oj.AddClass("isGood")
	fmt.Println(oj.Class()) // ["isGood"]
*/
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
Call ObjectDOM a function of Javascript

	//Example:
	oj.CallFunction("Remove",nil)
*/
func (o *ObjectDOM) CallFunction(function string, params ...interface{}) {
	o.object.Call(function, params)
}

/*
Set a property ObjectDOM property of Javascript

	//Example:
	oj.SetProperty("innerText","hello dev")
*/
func (o *ObjectDOM) SetProperty(property string, params ...interface{}) {
	o.object.Set(property, params)
}

// Getter a property as string
func (o *ObjectDOM) GetAsStringProperty(property string) string {
	return o.object.Get(property).String()
}

// Getter a property as js.Value
func (o *ObjectDOM) GetProperty(property string) js.Value {
	return o.object.Get(property)
}
