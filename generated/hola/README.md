# example-jsii-srcmak-go

> This example uses `patch-package` to patch jsii-srcmak to contain the wip Go support developed in this [PR](https://github.com/cdklabs/jsii-srcmak/pull/476).

This repo contains a simple Go program and uses JSII to include TypeScript code in the Go project.

It serves as an example to be able to talk about how Go support might work in [jsii-srcmak](https://github.com/cdklabs/jsii-srcmak).

## hello vs. hola
- `hello.ts` is included without a `go.mod` file and needs no require in the root `go.mod`
- `hola.ts` has a `go.mod` file and needs `require` & `replace` in the root `go.mod`.

If `hola` is used without `require`/`replace`, `go build` fails with the following error:
```
main.go:7:2: no required module provides package github.com/ansgarm/example-jsii-srcmak-go/generated/hola; to add it:
        go get github.com/ansgarm/example-jsii-srcmak-go/generated/hola
```

## Steps to run it
```
yarn install
yarn build
go run main.go
```
