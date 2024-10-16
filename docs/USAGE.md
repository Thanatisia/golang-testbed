# Documentation - Usage

## Synopsis/Syntax

```bash
testbench {optionals} <arguments> [actions ...]
```

## Parameters

### Positionals
- actions : Specify the positional keywords (aka 'actions') to execute
    - General
        + hello : Just prints out "world"
    - System Command Execution
        + execute-system-command-serial : Execute a specified System Command (by passing the `--set-cmd-str`) along with the System command arguments/parameters (by passing the `--set-cmd-args` parameters) and print the data out with communication
        + execute-system-command-realtime : Execute a specified System Command (by passing the `--set-cmd-str`) along with the System command arguments/parameters (by passing the `--set-cmd-args` parameters) and redirect/pipe the data stream into the standard input, output and error in real time
    - JSON Encoding/Parsing
        + print-json : Read the contents of a JSON file (currently hardcoded as 'test.json'), import into the application runtime, parse and uncompress/extract the JSON contents into a Dictionary/Key-Value Mapping of type <`map[string]<any>`>

### Optionals
- With Arguments
    > Note: You can separate the parameter/argument from the value either using the '=' delimiter or a space delimiter
    - System Command Execution
        - `--set-cmd-str{=}"your-command|executable|binary"` : Specify the command to execute
        - `--set-cmd-args{=}'arguments here ...'` : Specify the command arguments to pass into the command to be executed
    - `--set-value{=}'value'` : Test option to implement the receiving of values (with a delimiter)
- Flags
    + `-h | --help` : Displays help information
    + `-v | --version` : Display system version information
    + `-V | --verbose` : Enables Verbose display output
    + `-p | --print-all-arguments` : Print all arguments passed to the command line

## Usage

### Quickstart test
- Printing an hello world
    ```bash
    testbench hello
    ```

- Passing parameters
    - Using the `--set-value` test option to visualize
        - space delimited
            ```bash
            testbench --set-value "value"
            ```
        - '=' delimited
            ```bash
            testbench --set-value="value"
            ```

- Print verbose message
    ```bash
    testbench -V [actions]
    ```

### System Command Execution
- Execute command with communication and print to standard output
    - space delimited
        ```bash
        testbench --set-cmd-str "your-command-here" --set-cmd-args 'your arguments to pass here' execute-system-command-serial
        ```
    - '=' delimited
        ```bash
        testbench --set-cmd-str="your-command-here" --set-cmd-args='your arguments to pass here' execute-system-command-serial
        ```
    - Execute command with verbose output
        ```bash
        testbench -V --set-cmd-str{=}"your-command-here" --set-cmd-args{=}'your arguments to pass here' execute-system-command-serial
        ```

- Execute command with pipes/redirection and stream to standard output in real time
    - space delimited
        ```bash
        testbench --set-cmd-str "your-command-here" --set-cmd-args 'your arguments to pass here' execute-system-command-realtime
        ```
    - '=' delimited
        ```bash
        testbench --set-cmd-str="your-command-here" --set-cmd-args='your arguments to pass here' execute-system-command-realtime
        ```
    - Execute command with verbose output
        ```bash
        testbench -V --set-cmd-str{=}"your-command-here" --set-cmd-args{=}'your arguments to pass here' execute-system-command-realtime
        ```

