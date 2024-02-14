package dom

import (
	"fmt"
	"time"

	"github.com/Troxsoft/DisdyDOM/pkg"
)

func DomExample() {
	ele, err := pkg.Document.Create(pkg.H1)
	if err != nil {
		panic(err)
	}
	pkg.Document.AppendChild(ele)
	ele.SetValue("hola mundo")
	ele2, err := pkg.Document.Create(pkg.H1)
	if err != nil {
		panic(err)
	}
	pkg.Document.AppendChild(ele2)
	ele2.SetValue("hola mundo 2")
	style1 := ele2.NewStyle()
	style1.FromStylesSheet(pkg.CSStylesheet{
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

}
