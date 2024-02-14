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
	return &ObjectDOM{
		object: body,
	}
}
func (d *DocumentDOM) AppendChild(ele *ObjectDOM) {
	d.doc.Get("body").Call("appendChild", ele.object)
}
func (d *DocumentDOM) Create(typeOfElement int) (*ObjectDOM, error) {
	p := ConvertTypeHTMLToString(typeOfElement)
	if !IsValidTypeHTML(p) {
		return nil, ErrHTMLTypeInvalid
	}
	eleNew := d.doc.Call("createElement", p)
	return &ObjectDOM{
		object: eleNew,
	}, nil
}
