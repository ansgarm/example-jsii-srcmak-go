// 2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824
package hello

import (
	_init_ "github.com/ansgarm/example-jsii-srcmak-go/generated/hello/jsii"
	_jsii_ "github.com/aws/jsii-runtime-go"
)

type Hello interface {
	Hello() *string
}

// The jsii proxy struct for Hello
type jsiiProxy_Hello struct {
	_ byte // padding
}

func NewHello() Hello {
	_init_.Initialize()

	j := jsiiProxy_Hello{}

	_jsii_.Create(
		"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824.Hello",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewHello_Override(h Hello) {
	_init_.Initialize()

	_jsii_.Create(
		"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824.Hello",
		nil, // no parameters
		h,
	)
}

func (h *jsiiProxy_Hello) Hello() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"hello",
		nil, // no parameters
		&returns,
	)

	return returns
}

