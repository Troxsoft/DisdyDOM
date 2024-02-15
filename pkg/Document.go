package pkg

import "syscall/js"

//Represents a Document of Javascript
type DocumentDOM struct {
	doc js.Value
}

//Create a new DocumentDOM
//	//Example:
//	myDocument := NewDocumentDOM()
func NewDocumentDOM() *DocumentDOM {
	return &DocumentDOM{
		doc: js.Global().Get("document"),
	}
}

//Wrapper of document.getElementById() with alias
/*
	//Example:
	myInfo := Document.GetByID("myInfo")
*/
func (d *DocumentDOM) GetById(id string) *ObjectDOM {
	doc := js.Global().Get("document")
	body := doc.Call("getElementById", id)
	if body.IsNull() || body.IsUndefined() {
		return nil
	}
	p90 := ObjectDOM{
		object: body,
	}
	p90.Listener = NewEventListener(&p90)
	return &p90
}

//Wrapper of document.appendChild()
//	//Example:
//	Document.AppendChild(myObjectDOM)
func (d *DocumentDOM) AppendChild(ele *ObjectDOM) {
	d.doc.Get("body").Call("appendChild", ele.object)
}

// Wrapper of document.createElement()
//	//Example:
//	myInfo := Document.Create(H1)
func (d *DocumentDOM) Create(typeOfElement string) *ObjectDOM {
	eleNew := d.doc.Call("createElement", typeOfElement)
	P := &ObjectDOM{
		object: eleNew,
	}
	P.Listener = NewEventListener(P)
	return P
}
