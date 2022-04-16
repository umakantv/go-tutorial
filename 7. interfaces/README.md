# Go interfaces

In Go, we do not declare that some type implements an interface.  
Instead all types that satisfy the requirement for an interface implement it implicitly, which makes it a bit challenging as well to keep track of which types are implementing a partical interface.  

* Step 1:  
Make Concrete types implicitly implement some interface

* Step 2:  
Use that interface to reperesent those types for code reuse.

Interfaces only help in forming a contract between different concrete implementations.

Interfaces only support methods and other nested interfaces, since structs (custom types) can only attach methods, in absence of classes.

## Multiple Interfaces
Sometimes we may need to check whether a variable `v1` from interface type `I1` also supports interface type `I2`, in situations where we want to use a specific method (say `fooFromI2`) from `I2`.  
That can be done with following syntax:  
`v1WithFooFromI2, ok := v1.(I2)`  

If `ok` is `true`, `v1WithFooFromI2` is type-safe and will have all methods from `I2`.