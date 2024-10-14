// Define package name as this is the main entry point source file which the compiler will look for during compile-time
package main

// Import dependencies
import (
	// Built-in
	"fmt"
	"os"
	// "slices" // Built-in standard library package (introduced in August 2023) containing generic functionalities to handle slice types
	"strings"

	// Modules
	argparse "testbench_practiceground/cmd"
	// sqlite3db "testbench_practiceground/sqlite3db"
	// hello_world "testbench_practiceground/hello"
	syscallexec "testbench_practiceground/system_cmd_execution"
)

// Define main() function
func main() {
    // hello_world.Hello();
    // sqlite3db.ApplicationLogic();
    opts_Flags, opt_with_Arguments, positionals, error_msg := argparse.ParseArguments()

    // Error checking
    if len(error_msg) > 0 {
        fmt.Println("(-) Error encountered:")
        for i:=0; i < len(error_msg); i++ {
            fmt.Println(error_msg[i])
        }
    } else {
        // Iterate through the Key-Vaue Mappings (i.e. dictionary/map) using for-range

        if len(opts_Flags) > 0 {
            fmt.Println("")
            fmt.Println("Optional arguments - Flags")
            fmt.Println("")
            for key, value := range opts_Flags {
                fmt.Println("Argument [", key, "]: ", value)

                // Process the key-value mappings
                switch key {
                    case "print-all-arguments":
                        // Print all optional and positional arguments
                        fmt.Println("=====================")
                        fmt.Println("All Arguments Parsed:")
                        fmt.Println("=====================")
                        fmt.Println("Optional Flags:")
                        for key, value := range opts_Flags {
                            fmt.Println("- ", key, ":", value)
                        }
                        fmt.Println("Optional With Arguments:")
                        for key, value := range opt_with_Arguments {
                            fmt.Println("- ", key, ":", value)
                        }
                        fmt.Println("Positionals:")
                        for i:=0; i<len(positionals); i++ {
                            fmt.Println("- ", i, ":", positionals[i])
                        }
                        fmt.Println("==========")
                        fmt.Println(" LIST END ")
                        fmt.Println("==========")
                        os.Exit(0)
                    case "help":
                        fmt.Println("Display Help Message")
                    case "version":
                        fmt.Println("Display Application Version Information")
                    default:
                        fmt.Println("Invalid optional argument provided: ", key)
                }
            }
        } else {
            fmt.Println("")
            fmt.Println("No optional flags provided.")
            fmt.Println("")
        }

        if len(opt_with_Arguments) > 0 {
            fmt.Println("")
            fmt.Println("Optional arguments - With Values")
            fmt.Println("")
            for key, value := range opt_with_Arguments {
                fmt.Printf("Argument [%s]: %s\n", key, value)

                // Check if the key is in the argument list
                if _, is_found := opt_with_Arguments[key]; is_found == false {
                    fmt.Printf("Invalid optional argument provided: %s\n", key)
                } else {
                    // Process the key-value mappings
                    if key == "set-value" {
                        fmt.Println("Value Received: ", value)
                    }
                }
            }
        } else {
            fmt.Println("")
            fmt.Println("No optional arguments with values provided")
            fmt.Println("")
        }

        if len(positionals) > 0 {
            fmt.Println("")
            fmt.Println("Positional arguments")
            fmt.Println("")
            for i:=0; i < len(positionals); i++ {
                // Get current element
                var curr_element string = positionals[i]

                // Print current element
                fmt.Println(i, ":", curr_element)

                // Process the key-value mappings
                switch curr_element {
                    case "hello":
                        fmt.Println("World")
                    default:
                        fmt.Println("Invalid optional argument provided: ", curr_element)
                }
            }
        } else {
            fmt.Println("")
            fmt.Println("No positionals provided.")
            fmt.Println("")
        }
    }

    // Verify/validate verbose mode
    var verbose bool = opts_Flags["verbose"]

    fmt.Println("Verbose: ", verbose)

    // Check if "syscall-set-command" is in 'opt_with_Arguments' and Obtain the command line string to execute
    if cmd_str, is_found := opt_with_Arguments["syscall-set-command"]; is_found {
        fmt.Println("Command: ", cmd_str)

        // Check if 'syscall-set-arguments" is in 'opt_with_Arguments' and obtain the arguments
        if cmd_args, is_found := opt_with_Arguments["syscall-set-arguments"]; is_found {
            // Split the cmd_args string into a list
            var cmd_args_split []string = strings.Split(cmd_args, " ")

            fmt.Println("Arguments: ", cmd_args_split)

            // syscallexec.ExecSysCall(cmd_str, cmd_args_split...)
            syscallexec.ExecSysCallRealtime(cmd_str, verbose, cmd_args_split...)
            // syscallexec.ExecProcess(cmd_str, verbose, cmd_args_split...)
        } else {
            // syscallexec.ExecSysCall(cmd_str)
            syscallexec.ExecSysCallRealtime(cmd_str, verbose)
            // syscallexec.ExecProcess(cmd_str, verbose)
        }
    }
}

