# Re-Generating/Creation Steps

## Information

### Summary
+ This documentation contains steps taken to generate the baseline/root structure/hierarchy of the project (package) workspace

### Project Layout

```bash
[project-root-directory]/
|
|-- README.md
|-- go.mod
|-- main.go
|
|-- [module-name]/
    |
    |-- go.mod
    |-- *.go
    |
    |-- [submodule-name]/
        |
        |-- go.mod
        |-- *.go
```

## Setup

### Initialize a new working go project workspace (aka a new 'go package')

```bash
# Create a new go project workspace root directory (aka a new 'go package')
mkdir -pv /path/to/golang-testbed
cd /path/to/golang-testbed

# Initialize git version control system
git init
git config user.name [username]
git config user.email [email]
git checkout -b main
git remote add origin [remote-repository-server-url]

# Initialize a new 'go.mod' (go module definition file) for the package
go mod init golang_testbed
go mod tidy

# (Optional) Generate a template main.go entry point source file
cat <<EOF >> main.go
// Declare the 'main' special package
package main

// Import dependencies (libraries/packages/modules)
import (
    "fmt"
    "os"

    // To import a package with an alias name
    variable_alias "package/module/your-library-here"

    // To import a package from a git remote repository server (i.e. github)
    variable_alias "github.com/repo-author/repo-name"
)

// Define the main entry point function
func main() {
    fmt.Println("Hello World")
}
EOF
```

### Adding/Creating new local modules to the package

```bash
# Create a new module within the package
cd /path/to/golang-testbed
mkdir -pv ./path/to/module/name
cd ./path/to/module/name
go mod init golang_testbed/module-name
go mod tidy

# Map the new module within the package to the local relative path containing the module you created
cd /path/to/golang-testbed
echo -e "replace golang_testbed/module-name => ./path/to/module/name" >> go.mod
go mod tidy
```

### Adding/"Get"ting new external remote packages

```bash
# Get/pull/download the specified package from the remote repository server to the project local scope
cd /path/to/target/directory/go.mod
go get remote-repository-server-url
go mod tidy
```

### Adding/Creating new local modules to another module within the package

```bash
# Create a new module within the package
cd /path/to/golang-testbed
mkdir -pv ./path/to/module/name
cd ./path/to/module/name
go mod init golang_testbed/module-name
go mod tidy

# Create an indirect Map/Link of the module `[package-name]/[module-name]` into your project root directory's go.mod file
cd /path/to/target/module/go.mod
echo -e "replace golang_testbed/module-name => ./path/to/entry-point-module" >> go.mod
go mod tidy
```

## Documentations

## Snippets

## Wiki

### Terminologies
+ package: In golang, the package refers to the project workspace's root directory name
+ module: In golang, a module (or submodule) within a package refers to subdirector(ies) containing 1 or more `*.go` source files, all of which will provide functions, attributes/properties, variables and structures to the module namespace.

## Resources

## References

## Remarks

