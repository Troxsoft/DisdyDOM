package pkg

import (
	"fmt"
	"syscall/js"
)

// Create a rgb color (GENERATE A COLOR)
//
//	//Example:
//	myColor := CSSColor(255,0,0) // rgb(255,0,0)
func CSSColor(r, g, b int8) string {
	return fmt.Sprintf("rgb(%d,%d,%d)", r, g, b)
}

// Create a rgba color (GENERATE A COLOR)
//
//	//Example:
//	myColor := CSSColorWithAlpha(255,0,0,0.5) // rgba(255,0,0,0.5)
func CSSColorWithAlpha(r, g, b int8, a float32) string {
	return fmt.Sprintf("rgba(%d,%d,%d,%f)", r, g, b, a)
}

// Represents a Style.
// Represents a struct that can change the style of an element with various things
type CSStyle struct {
	ele   ObjectDOM
	style js.Value
}

// Inject css a DOM element
//
//	//Example:
//	myObjectDOM.NewStyle().Inject("color","blue")
func (style *CSStyle) Inject(key, value string) {
	style.style.Set(key, value)
}

//Injects CSS but with a structure. It can be useful for injecting various styles and for code autocompletion.
//	//Example:
/*	style1 := myObjectDOM.NewStyle()
	style1.FromStylesSheet(CSStylesheet{
		Color:           "red",
		BackgroundColor: "blue",
		TextAlign:       "center",
	})
*/
func (style *CSStyle) FromStylesSheet(sheet CSStylesheet) {
	sh := ConvertSheetToSlice(sheet)
	for key, value := range sh {

		style.Inject(key, value)
	}
}

// Create a new *CSStyle
func NewStyle(ele ObjectDOM) *CSStyle {
	return &CSStyle{
		ele:   ele,
		style: ele.object.Get("style"),
	}
}
