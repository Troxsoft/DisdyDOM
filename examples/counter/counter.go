package counter

import (
	"fmt"
	"time"

	"github.com/Troxsoft/DisdyDOM/pkg"
)

// GOARCH=wasm GOOS=js go build -o lib.wasm selectExample.go
func CounterExample() {
	co := 0
	obj := pkg.Document.GetById("counter")
	for {
		obj.SetValue(fmt.Sprint(co))
		time.Sleep(30 * time.Microsecond)
		co += 1
	}
}
