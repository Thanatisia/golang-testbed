# Quickstart Setup

## Setup

### Dependencies
+ go

### Pre-Requisites
- Set Environment Variables
    - `GOROOT=[go-sdk-directory]` : This environment variable is mapped to a custom directory containing your go (SDK) installation; You can find the go compiler, go tools or standard libraries in this directory; Optional - The Go binary assumes that go will be installed in `/usr/local/go` (in Linux) or `%HOMEDRIVE%\Go` (in Windows), but it is possible to install the Go tools to a different location
        - Windows
            ```dos
            SET GOROOT=%HOMEDRIVE%\go
            ```
        - Linux
            ```bash
            export GOROOT=/usr/local/go
            ```
    - `GOPATH=[your-go-home-directory-here]` : Specify your Go home/project workspace here (similar to the 'PREFIX' environment variable for C); The packages/libraries/modules installed via `go install` will be installed into the package folder here here
        - Windows
            + Default: `%USERPROFILE%\go`
            ```dos
            SET GOPATH=%USERPROFILE%\go
            ```
        - Linux
            + Default: `$HOME/go`
            ```bash
            export GOPATH=$HOME/go
            ```
    - `GOBIN=[your-go-binary-directory-here]` : Specify your Go binaries directory here; This environment variable is mapped to the path where the binaries installed via `go install` will be installed into
        - Windows
            ```dos
            SET GOBIN=%GOROOT%\bin
            ```
        - Linux
            ```bash
            export GOBIN=$GOROOT/bin
            ```

- Install Golang
    - Using Package Manager
        - apt-based (Debian)
            ```bash
            apt install golang
            ```
        - pacman-based (ArchLinux)
            ```bash
            pacman -S golang
            ```

- (Optional) Append the golang directory into the PATH system environment variable
    - Programmatically using the PATH environment variable
        - Bash
            ```bash
            export PATH=/path/to/go/bin:$PATH
            ```
        - Batch
            ```dos
            SET PATH=\path\to\go\bin;%PATH%
            ```
    - Manually
        - Open 'System Properties'
            - Press the button 'Environment Variables'
                - Under 'System Variables', double click 'PATH' to modify the system environment variable path
                    + Press 'New' and type the path to the bin directory (i.e. \path\to\go\bin)
                    + Press 'OK' to confirm
                + Press 'OK' to apply
            + Press 'Apply' and 'OK' to close

### Project workspace structure/layout setup - main.go in root directory

#### Project Layout

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

#### Quickstart Setup

> Initialize a new working go package (project)

- Create a new go project workspace root directory (aka a new 'go package')
    ```bash
    mkdir -pv /path/to/project-root-directory
    cd /path/to/project-root-directory
    ```

- Initialize git version control system
    ```bash
    git init
    git config user.name [username]
    git config user.email [email]
    git checkout -b main
    git remote add origin [remote-repository-server-url]
    ```

- Initialize a new 'go.mod' (go module definition file) for the package
    - Explanation
        + This file defines the project's module repository and all external dependencies required by the project
    ```bash
    go mod init project-root-directory-name
    go mod tidy
    ```

- (Optional) Generate a template main.go entry point source file
    ```bash
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

> Adding/Creating new local modules to the package

- Create a new module within the package
    ```bash
    cd /path/to/project-root-directory
    mkdir -pv ./path/to/module/name
    cd ./path/to/module/name
    go mod init project-root-directory-name/module-name
    go mod tidy
    ```

- Map the new module within the package to the local relative path containing the module you created
    ```bash
    cd /path/to/project-root-directory
    echo -e "replace project-root-directory-name/module-name => ./path/to/module/name" >> go.mod
    go mod tidy
    ```

> Adding/"Get"ting new external remote packages

- Get/pull/download the specified package from the remote repository server to the project local scope
    ```bash
    cd /path/to/target/directory/go.mod
    go get remote-repository-server-url
    go mod tidy
    ```

> Adding/Creating new local modules to another module within the package

- Create a new module within the package
    ```bash
    cd /path/to/project-root-directory
    mkdir -pv ./path/to/module/name
    cd ./path/to/module/name
    go mod init project-root-directory-name/module-name
    go mod tidy
    ```

- Create an indirect Map/Link of the module `[package-name]/[module-name]` into your project root directory's go.mod file
    ```bash
    cd /path/to/target/module/go.mod
    echo -e "replace project-root-directory-name/module-name => ./path/to/entry-point-module" >> go.mod
    go mod tidy
    ```

#### Snippets and examples

> Project workspace is named 'package'

- Information and Properties
    + Project workspace directory: package
    + Package Name: package
    + Entry Point: ./main.go
    - Modules
        + New Module: ./modules/new_module

- Steps
    - Setup go package
        ```bash
        mkdir -pv package/
        cd package/
        go mod init package
        go mod tidy
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
    - Setup subdirector(ies) for the modules in the package
        ```bash
        mkdir -pv modules/new_module
        cd modules/new_module
        go mod init package/new_module
        go mod tidy
        ```
    - Map the module namespace to the module directory's path in the project workspace
        ```bash
        cd ../..
        echo -e "replace package/new_module => ./modules/new_module" >> go.mod
        go mod tidy
        ```
    - Build/Compile the binary/executable
        ```bash
        go build -o dist/bin/output .
        ```
    - Run the package
        ```bash
        go run .
        ```

### Project workspace structure/layout setup - main.go in module subdirectory

#### Project Layout

```bash
[project-root-directory]/
|
|-- README.md
|-- go.mod
|
|-- [entry-point-module]/
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

#### Quickstart Setup

> Initialize a new working go project workspace (aka a new 'go package')

- Create a new go project workspace root directory (aka a new 'go package')
    ```bash
    mkdir -pv /path/to/project-root-directory
    cd /path/to/project-root-directory
    ```

- Initialize git version control system
    ```bash
    git init
    git config user.name [username]
    git config user.email [email]
    git checkout -b main
    git remote add origin [remote-repository-server-url]
    ```

- Initialize a new 'go.mod' (go module definition file) for the package
    ```bash
    go mod init project-root-directory-name
    go mod tidy
    ```

- Create a new module directory for the 'main.go' entry point source codes
    ```bash
    mkdir -pv ./path/to/entry-point-module/
    cd ./path/to/entry-point-module/
    go mod init project-root-directory-name/entry-point-module
    go mod tidy
    ```

- (Optional) Generate a template main.go entry point source file
    ```bash
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

- Map the new (entry point) module within the package to the local relative path containing the module you created
    ```bash
    cd /path/to/project-root-directory
    echo -e "replace project-root-directory-name/entry-point-module => ./path/to/entry-point-module" >> go.mod
    go mod tidy
    ```

> Adding/Creating new local modules to the package

- Create a new module within the package
    ```bash
    cd /path/to/project-root-directory
    mkdir -pv ./path/to/module/name
    cd ./path/to/module/name
    go mod init project-root-directory-name/module-name
    go mod tidy
    ```

- Create an indirect Map/Link of the module `[package-name]/[module-name]` into your project root directory's go.mod file
    ```bash
    cd /path/to/project-root-directory
    go get project-root-directory-name/module-name
    go mod tidy
    ```

> Adding/"Get"ting new external remote packages

- Get/pull/download the specified package from the remote repository server to the project local scope
    ```bash
    cd /path/to/target/directory/go.mod
    go get remote-repository-server-url
    go mod tidy
    ```

> Adding/Creating new local modules to another module within the package

- Create a new module within the package
    ```bash
    cd /path/to/project-root-directory
    mkdir -pv ./path/to/module/name
    cd ./path/to/module/name
    go mod init project-root-directory-name/module-name
    go mod tidy
    ```

- Create an indirect Map/Link of the module `[package-name]/[module-name]` into your project root directory's go.mod file
    ```bash
    cd /path/to/target/module/go.mod
    echo -e "replace project-root-directory-name/module-name => ./path/to/entry-point-module" >> go.mod
    go mod tidy
    ```

#### Snippets and examples

> Project workspace is named 'package'

- Information and Properties
    + Project workspace directory: package
    + Package Name: package
    - Modules
        + Entry Point: modules/main

- Steps
    - Setup go package
        ```bash
        mkdir -pv package/
        cd package/
        go mod init package
        go mod tidy
        ```
    - Setup module subdirectory for the main entry point source codes
        ```bash
        mkdir -pv modules/main
        cd modules/main
        go mod init package/main
        go mod tidy
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
    - Map the main module to the module directory's path in the project workspace
        ```bash
        cd ../..
        echo -e "replace package/main => ./modules/main" >> go.mod
        go mod tidy
        ```
    - Build/Compile the binary/executable
        ```bash
        go build -o dist/bin/output modules/main/main.go
        ```
    - Run the package
        ```bash
        go run modules/main/main.go
        ```

## Resources

## References

## Remarks

