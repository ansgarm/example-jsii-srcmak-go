module github.com/ansgarm/example-jsii-srcmak-go

go 1.16

require (
    github.com/aws/jsii-runtime-go v1.29.0
    github.com/ansgarm/example-jsii-srcmak-go/generated/hola v0.0.0
)

replace github.com/ansgarm/example-jsii-srcmak-go/generated/hola => ./generated/hola