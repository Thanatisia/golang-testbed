# CHANGELOGS

## Table of Contents
+ [2024-10-14](#2024-10-14)
+ [2024-10-15](#2024-10-15)
+ [2024-10-16](#2024-10-16)
+ [2024-10-23](#2024-10-23)

## Entries
### 2024-10-14
#### 1629H
+ Initial Commit

- New
    + Added new document 'README.md'
    + Added new document 'CHANGELOGS.md'
    + Added new document 'GENERATING.md'
    + Added new document 'Makefile'
    + Added new golang project (package) module definition file 'go.mod'
    + Added new golang source file 'main.go'
    - Added new directory 'docs/'
        - Added new directory 'snippets/'
            + Added new golang snippet 'cli-argument-parsing-delimiter-parameters.go'
            + Added new golang snippet 'cli-argument-parsing-manual.go'
    - Added new directory 'src/'
        - Added new golang package module directory 'cmd' for Command Prompt/CLI-related functionalities
            + Added new golang source file 'cli.go'
            + Added new golang module definition file 'go.mod'
        - Added new directory 'modules'
            - Added new directory 'features'
                - Added new module directory 'sqlite3db' for feature 'SQLite3'
                    + Added new golang module definition file 'go.mod'
                    + Added new golang source file 'app.go' for the Application UI and Implementation
                    + Added new golang source file 'sqlite3db.go' for the Core Database Logic Library
                - Added new module directory 'system_cmd_execution' for feature 'System Command Execution' and 'Subprocess Calls'
                    + Added new golang module definition file 'go.mod'
                    + Added new golang source file 'syscall.go' for functionalities regarding System Call and Command Execution using 'syscall'
            - Added new directory 'tutorials'
                - Added new module directory 'hello' for Tutorial project 'Hello World'
                    + Added new golang module definition file 'go.mod'
                    + Added new golang source file 'hello.go'
    - Added new directory 'tests/' to be a test package
        + Added new document 'go.Makefile'
        + Added new test golang project (package) module definition file 'go.mod'
        + Added new golang source file 'main.go'
        - Added new directory 'src'
            - Added new module directory 'cli'
                + Added new golang module definition file 'go.mod'
                + Added new test module source file 'parser.go'
                + Added new test module source file 'test.go' for testing (Unused)
            - Added new module directory 'cmd'
                + Added new golang module definition file 'go.mod'
                + Added new test module source file 'test.go' for testing (Unused)

#### 1633H
- New
    - Added new document 'BUILD.md' in 'docs/' to perform a quickstart setup working golang project (package) workspace structure/hierarchy

- Updates
    - Migrated 'GENERATING.md' to 'docs/'
        + Added setup information

#### 1647H
- Updates
    - Migrated 'BUILD.md' to 'docs/'
        + Fixed bash shellscript block and cat-EOF output

#### 1701H
- Updates
    - Updated document 'GENERATING.md' in 'docs/'
        + Renamed package from 'testbench_practiceground' => 'golang_testbed'
    - Updated package module definition file 'go.mod'
        + Renamed package from 'testbench_practiceground' => 'golang_testbed'
    - Updated golang entry point source file 'main.go'
        + Renamed package and modules from 'testbench_practiceground' => 'golang_testbed'
    - Updated module definition file 'go.mod' in 'src/cmd/'
        + Renamed package and modules from 'testbench_practiceground' => 'golang_testbed'
    - Updated module definition file 'go.mod' in 'src/modules/features/sqlite3db/'
        + Renamed package and modules from 'testbench_practiceground' => 'golang_testbed'
    - Updated module definition file 'go.mod' in 'src/modules/features/system_cmd_execution/'
        + Renamed package and modules from 'testbench_practiceground' => 'golang_testbed'
    - Updated module definition file 'go.mod' in 'src/modules/tutorials/hello/'
        + Renamed package and modules from 'testbench_practiceground' => 'golang_testbed'

### 2024-10-15
#### 1014H
- New
    - Added new module directory 'jsonio' in 'src/modules/features/' to test JSON Encoding/Parsing
        - Added new go module definition file 'go.mod'
        - Added new golang module library file 'jsonio.go'
- Updates
    - Updated golang package module definition file 'go.mod'
        + Added 'replace' alias for new module 'jsonio'
        + go mod tidy
    - Updated golang source file 'main.go'
        + Added new import for jsonio
        + Added a new optional switch case for verbose mode
        + Added new positional keywords
    - Updated golang source file 'cli.go' in 'src/cmd'
        + Fixed bug where separating parameters with a space delimiter (' ') doesnt keep the value

### 2024-10-16
#### 1726H
- New
    + Added new document 'USAGE.md' in 'docs/' for Post-setup/build usage
- Updates
    - Update document 'README.md'
        + Added installation and uninstallation steps
    - Updated golang source file 'main.go'
        + Updated positional action's keyword identifier

### 2024-10-23
#### 2212H
- New
    - Added new document 'README.md' in 'src/modules/features/sqlite3db' for the implementation explanation of sqlite3
    - Added new golang source library 'todo.go' in module 'src/modules/features/sqlite3db/'
    - Added new directory 'tests' in sqlite3db for housing 'app.go' and 'sqlite3db.go' temporarily
- Updates
    - Updated golang package module definition file 'go.mod'
        + Added package local module 'sqlite3db' as dependency
    - Updated golang source file 'main.go'
        + Uncommented 'golang_testbed/sqlite3db' local module dependency
        + Added new positional argument 'start-todolist-webserver'
    - Updated golang module definition file 'go.mod' in 'src/modules/features/sqlite3db/'
        + Updated version of external module 'mattn/go-sqlite3'
    - Migrated 'app.go' and 'sqlite3db.go' in 'src/modules/features/sqlite3db' => 'src/modules/features/sqlite3db/tests' temporarily



