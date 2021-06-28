- When we need to call code that may have been written by someone else,
we can look for a package that has functions we can use in our code.

- pkg.go.dev site is a very good place to find packages we need.

- When we want golang to add packages imported as requirement, as well as a go.sum file for use in authenticating the module.
    run the command: ``` go mod tidy ```
    
Exercise:
1. Go to pkg.go.dev and search for quote package. Locate and click the rsc.io/quote package in search results.
    then import rsc.io/quote in your program.