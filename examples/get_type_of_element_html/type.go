package gettypeofelementhtml

import (
	"fmt"

	"github.com/Troxsoft/DisdyDOM/pkg"
)

func PrintTypeOfElemnt(element pkg.ObjectDOM) {
	fmt.Println(element.GetHTMLType())
}
