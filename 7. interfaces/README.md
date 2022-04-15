# Go interfaces

In Go, we do not declare that some type implements an intercace.  
Instead all types that satisfy the requirement for an interface implement it implicitly, which makes it a bit challenging as well to keep track of which types are implementing a partical interface.  

* Step 1:  
Make Concrete types implicitly implement some interface

* Step 2:  
Use that interface to reperesent those types for code reuse.