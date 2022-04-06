# creating a module

https://go.dev/doc/tutorial/create-module.html

A multi-part tutorial that introduces common programming language features from the Go perspective.

Go modules are a collection of packages.  
> In a module, you collect one or more related packages for a discrete and useful set of functions.

A package is used to collect related functions.

And in that sense, your code is also a module that uses other modules.  

A module specifies the Go version it uses and its dependencies.
* We will make a module named `greetings` here.
* We start a module by the `go mod init <module_name>` command.
* We declare the package for each go file such as `greetings.go` declares that it is a part of `greetings` package.

In Go, a function whose name starts with a capital letter can be called by a function not in the same package.