package main

import (
	//gettypeofelementhtml "github.com/Troxsoft/DisdyDOM/examples/get_type_of_element_html"
	dom "github.com/Troxsoft/DisdyDOM/examples/Dom"
	_ "github.com/Troxsoft/DisdyDOM/pkg"
)

func main() {
	c := make(chan struct{}, 0)
	dom.DomExample()

	//gettypeofelementhtml.PrintTypeOfElemnt(*pkg.Document.GetById("hello"))
	<-c
}
