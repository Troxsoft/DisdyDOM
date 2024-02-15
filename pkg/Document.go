package pkg

import "syscall/js"

type DocumentDOM struct {
	doc js.Value
}

func NewDocumentDOM() *DocumentDOM {
	return &DocumentDOM{
		doc: js.Global().Get("document"),
	}
}

// document.getElementById
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
func (d *DocumentDOM) AppendChild(ele *ObjectDOM) {
	d.doc.Get("body").Call("appendChild", ele.object)
}
func (d *DocumentDOM) Create(typeOfElement string) *ObjectDOM {
	eleNew := d.doc.Call("createElement", typeOfElement)
	P := &ObjectDOM{
		object: eleNew,
	}
	P.Listener = NewEventListener(P)
	return P
}
