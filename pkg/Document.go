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
