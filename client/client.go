package main

import (
	"fmt"
	"strings"
	"syscall/js"
)

func main() {
	wasmApp := js.Global().Get("document").Call("getElementById", "go_wasm")
	wasmApp.Set("innerHTML", "<h1> Golang WASM TODO app </h1>")

	fmt.Println("Hello World!")
}

// type DomElement interface {
// 	ToHTML() string
// }

// type baseElement struct {
// 	children []*BaseElement
// 	class    []string
// 	tag      string
// 	wasmID   string
// 	parentID string
// }

// func (b *baseElement) ToHTML() string {
// 	classStr := strings.Join(be.class[:], " ")
// 	childrenStr := b.GenerateChildTree()
// 	return fmt.Sprintf("<%s class=\"%s\" wasm_id=\"%s\"></%s>", b.tag, classStr, b.wasmID, b.tag)
// }

// func (b *baseElement) GenerateChildTree() {
// 	for c := range b.children {
// 		*c.ToHTML()
// 	}
// }

// func (b Banner) ToHTML(msg) string {
// 	return "<h3 class=\"inverted ui block header\"> Todo WASM </h3>"
// }
