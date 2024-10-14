# CHANGELOGS

## Table of Contents
+ [2024-10-14](#2024-10-14)

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

