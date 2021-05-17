# example-jsii-srcmak-go

> This example uses `patch-package` to patch jsii-srcmak to contain the wip Go support developed in this [PR](https://github.com/cdklabs/jsii-srcmak/pull/476).

This repo contains a simple Go program and uses JSII to include TypeScript code in the Go project.  

It serves as an example to be able to talk about how Go support might work in [jsii-srcmak](https://github.com/cdklabs/jsii-srcmak).

**The generated code is checked in to be able to link to it**

## Steps to run it
```
yarn install
yarn build
go run main.go # should print "Hello World\nHola\n"
```