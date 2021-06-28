- Golang program imports packages that contain other modules.
- We manage those dependencies through our own module.
- That module is defined by a go.mod file that tracks the modules that provide those packages.
- That go.mod file stays with our code, including in our source code repository.

To enable dependency tracking for our code by creating a go.mod file, run the ```go mod init command```,
giving it the name of the module our code will be in. The name is the module's module path,
this will be the repository location where your source code will be kept, such as github.com/mymodule.
If you plan to publish your module for others to use, the module path must be a location from which Go tools can download your module.

Exercise:
1. Create a module, enable dependencies tracking.
$ ``` go mod init learn-golang/'1.getting_started' ```