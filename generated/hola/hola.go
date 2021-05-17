// b221d9dbb083a7f33428d7c2a3c3198ae925614d70210e28716ccaa7cd4ddb79
package hola

import (
	_init_ "github.com/ansgarm/hola-go/hola/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"
)

type Hola interface {
	Hola() *string
}

// The jsii proxy struct for Hola
type jsiiProxy_Hola struct {
	_ byte // padding
}

func NewHola() Hola {
	_init_.Initialize()

	j := jsiiProxy_Hola{}

	_jsii_.Create(
		"b221d9dbb083a7f33428d7c2a3c3198ae925614d70210e28716ccaa7cd4ddb79.Hola",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewHola_Override(h Hola) {
	_init_.Initialize()

	_jsii_.Create(
		"b221d9dbb083a7f33428d7c2a3c3198ae925614d70210e28716ccaa7cd4ddb79.Hola",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_Hola) Hola() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"hola",
		nil, // no parameters
		&returns,
	)

	return returns
}

