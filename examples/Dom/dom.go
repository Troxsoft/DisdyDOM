package dom

import (
	"fmt"
	"time"

	. "github.com/Troxsoft/DisdyDOM/pkg"
)

func DomExample() {
	ele := Document.Create(H1)

	ExportFunction("saludamePapi", func(f []any) any {
		Window.Alert(fmt.Sprintf("hola como estas: %s", f[0]))
		return fmt.Sprintf("hola como estas: %s", f[0])
	})
	Document.AppendChild(ele)
	ele.SetValue("hola mundo")
	ele2 := Document.Create(P)

	Document.AppendChild(ele2)
	ele2.SetValue("hola mundo 2")
	style1 := ele2.NewStyle()
	style1.FromStylesSheet(CSStylesheet{
		Color:           "red",
		BackgroundColor: "blue",
		TextAlign:       "center",
	})
	time.Sleep(2 * time.Second)
	ele2.AddClass("pl")
	ele2.SetID("colombia")
	ele2.AddClass("pl2")
	fmt.Println(ele2.Class())
	fmt.Println(ele2.ID())

	time.Sleep(2 * time.Second)
	ele2.ClearID()
	time.Sleep(1 * time.Second)
	ele2.SetProperty("id", "colombia")
	Window.Alert(ele2.ID())
	button := Document.GetById("btn1")
	button.SetValue("click")
	button.Listener.AddClick("mostrar alerta", func(me MouseEvent) {
		fmt.Println(button.Listener.GetClicks())
		Window.Alert(fmt.Sprintf("x:%d | y:%d", me.X, me.Y))
	})

}
