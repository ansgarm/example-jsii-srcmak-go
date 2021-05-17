package hello

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go"
)

func init() {
	_jsii_.RegisterClass(
		"2cf24dba5fb0a30e26e83b2ac5b9e29e1b161e5c1fa7425e73043362938b9824.Hello",
		reflect.TypeOf((*Hello)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "hello", GoMethod: "Hello"},
		},
		func() interface{} {
			return &jsiiProxy_Hello{}
		},
	)
}
