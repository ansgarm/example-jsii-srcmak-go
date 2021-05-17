package main

import (
	"fmt"

	hello "github.com/ansgarm/example-jsii-srcmak-go/generated/hello"
	hola "github.com/ansgarm/hola-go/hola"
)

func main() {
	fmt.Printf("%v\n", *hello.NewHello().Hello())
	fmt.Printf("%v\n", *hola.NewHola().Hola())
}
