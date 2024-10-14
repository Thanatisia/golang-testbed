# Go Practice Project

## Information
### Description

This is a simple Testbench/Practice ground project in which, will be used to test and implement various use cases, backend/frontend functions, UI components, modules, class, packages for reference.

The contents may also become a massive GUI with buttons to visualize all the features in a safe, local development environment.

## Setup

### Dependencies
+ go

### Pre-Requisites
- Set Environment Variables
    + GOROOT
    + GOPATH
    + GOBIN

- Create project workspace director(ies)
    ```bash
    mkdir -pv dist/bin
    ```

### Build/Compile

> Using go
- Compile/Build the project locally
    ```bash
    go build .
    ```

- Compile and output executable/binary with a custom file name
    ```bash
    go build {-o [custom-executable-filename]} [workspace-directory|.]
    ```

- Compile and output executable/binary with a custom output directory and file name
    ```bash
    go build {-o dist/bin/[custom-executable-filename]} [workspace-directory|.]
    ```

> Using Makefile
- Compile/Build the project locally
    ```bash
    make GO_FLAGS="your go flags here" OUT_PATH="your-custom-output-directory-filepath" build
    ```

- Compile/Build and Run the executable, then delete the executable
    ```bash
    make GO_RUN=[path-to-compile-and-run (default: .)] GO_RUN_FLAGS="flags-to-pass/parse-to-application-CLI-argument" run
    ```

- Get/Install a new go module/library/package from a remote repository server
    ```bash
    make PKG_DEPS="your package/module/library dependency URLs here" pull
    ```

- Append a new entry mapping your target local module's alias/name to its path within the project workspace structure/hierarchy (`replace [module path alias] => [module path]`) and writing into the go.mod file
    ```bash
    make MOD_NAME=testbench_practiceground/sqlite3db MOD_PATH=modules/features/sqlite3db append-module
    ```

- Install all specified dependencies in the go repository module definitions file ('go.mod') using `go mod tidy`
    ```bash
    make refresh
    ```

### Installation
- Install the project to the go binary directory (`<GOBIN>|<GOPATH>/bin`)
    ```bash
    go install .
    ```

### Uninstallation
- Uninstall the binary from the go binary directory (`<GOBIN>|<GOPATH>/bin`)
    ```bash
    rm [GOBIN]/executable-name
    ```

## Documentations

### Table of Content and Topics
+ [Makefile](Makefile)
+ [Main Entry Point and CLI Argument Usage](main.go)
+ [CLI Argument Parsing](src/cmd/cli.go)
- Features
    - SQLite3
        + [SQLite3 Database Test Module - Application UI and Implementation](src/modules/features/sqlite3db/app.go)
        + [SQLite3 Database Test Module - Core Database Logic Library](src/modules/features/sqlite3db/sqlite3db.go)
    - System Call and Command Execution
        + [System Call and Command Execution - Using syscall](src/modules/features/system_cmd_execution/syscall.go)
- Tutorials
    - [Hello World](src/modules/tutorials/hello/hello.go)
- Proof of Concept and DevOps (Development & Operations Idea) Testbench
    + [Main Entry Point](tests/main.go)
    + [CLI Argument Parser](tests/src/cli/parser.go)
    + [Package Module (tests/cli) Test File (Unused)](tests/src/cli/test.go)
    + [Package Module (tests/cmd) Test File (Unused)](tests/src/cmd/test.go)
- Documentations
    - Snippets
        + [Manual/Generic CLI Argument Parsing - Without Delimiter-separated Parameters](docs/snippets/cli-argument-parsing-manual.go)
        + [Manual/Generic CLI Argument Parsing - With Delimiter-separated Parameters](docs/snippets/cli-argument-parsing-delimiter-parameters.go)
    - Templates

### Synopsis/Syntax

### Parameters

### Usage

## Wiki

### Steps
- To initialize repository as a go module
    ```bash
    go mod init [repository-url]
    ```

## Resources

## References

## Remarks

