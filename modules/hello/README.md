This is the module where we import and use the earlier created `greetings` module.

## Initialize module
To enable dependency tracking for our code, we initiate `go.mod` file by `go mod init <module_name>`

## Adding dependencies
We create dependencies when you imported the greetings package in a go code file.

## Importing dependencies in modules: 

We would usually use packages from remote repositories, so we would require them from their location.

But here we need to add local dependency, so we tell go compiler where to look for it with this command `go mod edit -replace greetings=../greetings`. This would change the go.mod file as follows: 

```diff
module hello

go 1.18

+ replace greetings => ../greetings
```

## Using `require`
Then we run `go mod tidy` to update the `go.mod` file with the `require` statement.

That's it! We are done. We can now run the module with `go run .`.