package pkg

import (
	"fmt"
	"strings"
	"syscall/js"
)

// HTML ELEMENT
const (
	DOCTYPE    = "!DOCTYPE"
	HTML       = "HTML"
	HEAD       = "HEAD"
	META       = "META"
	TITLE      = "TITLE"
	LINK       = "LINK"
	STYLE      = "STYLE"
	BODY       = "BODY"
	H1         = "H1"
	H2         = "H2"
	H3         = "H3"
	H4         = "H4"
	H5         = "H5"
	H6         = "H6"
	P          = "P"
	BR         = "BR"
	HR         = "HR"
	A          = "A"
	EM         = "EM"
	STRONG     = "STRONG"
	SPAN       = "SPAN"
	DIV        = "DIV"
	IMG        = "IMG"
	UL         = "UL"
	OL         = "OL"
	LI         = "LI"
	TABLE      = "TABLE"
	TR         = "TR"
	TH         = "TH"
	TD         = "TD"
	CAPTION    = "CAPTION"
	THEAD      = "THEAD"
	TBODY      = "TBODY"
	TFOOT      = "TFOOT"
	FORM       = "FORM"
	INPUT      = "INPUT"
	BUTTON     = "BUTTON"
	SELECT     = "SELECT"
	OPTION     = "OPTION"
	TEXTAREA   = "TEXTAREA"
	LABEL      = "LABEL"
	FIELDSET   = "FIELDSET"
	LEGEND     = "LEGEND"
	SCRIPT     = "SCRIPT"
	NOSCRIPT   = "NOSCRIPT"
	AUDIO      = "AUDIO"
	VIDEO      = "VIDEO"
	SOURCE     = "SOURCE"
	CANVAS     = "CANVAS"
	EMBED      = "EMBED"
	IFRAME     = "IFRAME"
	OBJECT     = "OBJECT"
	PARAM      = "PARAM"
	SVG        = "SVG"
	MATH       = "MATH"
	DEL        = "DEL"
	INS        = "INS"
	CODE       = "CODE"
	KBD        = "KBD"
	SAMP       = "SAMP"
	VAR        = "VAR"
	SUB        = "SUB"
	SUP        = "SUP"
	Q          = "Q"
	BLOCKQUOTE = "BLOCKQUOTE"
	MARK       = "MARK"
	TIME       = "TIME"
	PROGRESS   = "PROGRESS"
	METER      = "METER"
	I          = "I"
	B          = "B"
	U          = "U"
	S          = "S"
	SMALL      = "SMALL"
	STRIKE     = "STRIKE"
	BIG        = "BIG"
	FONT       = "FONT"
	CENTER     = "CENTER"
	PRE        = "PRE"
	ARTICLE    = "ARTICLE"
	ASIDE      = "ASIDE"
	DETAILS    = "DETAILS"
	FIGURE     = "FIGURE"
	FIGCAPTION = "FIGCAPTION"
	FOOTER     = "FOOTER"
	HEADER     = "HEADER"
	MAIN       = "MAIN"
	MENU       = "MENU"
	NAV        = "NAV"
	SECTION    = "SECTION"
	SUMMARY    = "SUMMARY"
)

// A *DocumentDOM
var Document = NewDocumentDOM()

// A *WindowDOM
var Window = NewWindowDOM()

/*
Export a function to javascipt. You can call it by name and pass parameters and return them

	//Example:
	ExportFunction("saludamePapi", func(f []any) any {
		Window.Alert(fmt.Sprintf("hola como estas: %s", f[0]))
		return fmt.Sprintf("hola como estas: %s", f[0])
	})

To call in javascript

		const go = new Go();

	        WebAssembly.instantiateStreaming(fetch("lib.wasm"), go.importObject).then((result) => {
	            go.run(result.instance);
	            let h = saludamePapi("juan")
	            console.table({ saludo: h })
	        });
*/
func ExportFunction(funcName string, function func(f []any) any) {

	js.Global().Set(funcName, js.FuncOf(func(this js.Value, inputs []js.Value) any {
		params := []any{}
		for _, value := range inputs {
			if value.Type() == js.TypeString {
				params = append(params, value.String())
			} else if value.Type() == js.TypeBoolean {
				params = append(params, value.Bool())
			} else if value.Type() == js.TypeNumber {
				q := value.Float()
				if strings.Contains(fmt.Sprint(q), ".") {

					params = append(params, value.Float())
				} else {
					params = append(params, value.Int())
				}
			}
		}
		return function(params)
	}))
}
