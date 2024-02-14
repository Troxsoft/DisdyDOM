package pkg

import (
	"syscall/js"
)

// some colors
var (
	RED   = NewRGB(255, 0, 0)
	BLUE  = NewRGB(0, 0, 255)
	GREEN = NewRGB(0, 255, 0)
)

type RGBColor struct {
	R uint8
	G uint8
	B uint8
	A float32
}

func NewRGB(r, g, b uint8) RGBColor {
	return RGBColor{
		R: r,
		G: g,
		B: b,
		A: 1.0,
	}
}
func NewRGA(r, g, b uint8, a float32) RGBColor {
	return RGBColor{
		R: r,
		G: g,
		B: b,
		A: a,
	}
}

type CSStyle struct {
	ele   ObjectDOM
	style js.Value
}

func (style *CSStyle) Inject(key, value string) {
	style.style.Set(key, value)
}
func (style *CSStyle) FromStylesSheet(sheet CSStylesheet) {
	sh := ConvertSheetToSlice(sheet)
	for key, value := range sh {

		style.Inject(key, value)
	}
}

func NewStyle(ele ObjectDOM) *CSStyle {
	return &CSStyle{
		ele:   ele,
		style: ele.object.Get("style"),
	}
}
