# Workspaces
[Multi-modules workspace](https://go.dev/doc/tutorial/workspaces)

1. Make a folder named `hello`. CD into it, create module using `go mod init example.com/hello`. This will be treated as the main module.
2. Initiate a workspace using `go work init .hello`. This will create a `go.work` file.
3. Now go to the workspace directory, and clone `https://go.googlesource.com/example` as another module. From this module, we will use `stringutils` package.
4. Now add ./example as a module to the workspace with `go work use ./example`.
5. go.work can be used instead of adding `replace` directives to work across multiple modules.
6. Now we can add our own functionality to `./example` module and use it in our code.