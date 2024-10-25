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
	argparse "golang_testbed/cmd"
	sqlite3db "golang_testbed/sqlite3db"
	// hello_world "testbench_practiceground/hello"
	syscallexec "golang_testbed/system_cmd_execution"
    "golang_testbed/jsonio"
)

// Define main() function
func main() {
    // hello_world.Hello();
    // sqlite3db.ApplicationLogic();
    opts_Flags, opt_with_Arguments, positionals, error_msg := argparse.ParseArguments()

    // Initialize Variables
    var verbose bool = false

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
                    case "verbose":
                        // Verify/validate verbose mode
                        verbose = opts_Flags["verbose"]
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
                    case "execute-system-command-serial":
                        fmt.Println("Verbose: ", verbose)

                        // Check if "syscall-set-command" is in 'opt_with_Arguments' and Obtain the command line string to execute
                        if cmd_str, is_found := opt_with_Arguments["syscall-set-command"]; is_found {
                            fmt.Println("Command: ", cmd_str)

                            // Check if 'syscall-set-arguments" is in 'opt_with_Arguments' and obtain the arguments
                            if cmd_args, is_found := opt_with_Arguments["syscall-set-arguments"]; is_found {
                                // Split the cmd_args string into a list
                                var cmd_args_split []string = strings.Split(cmd_args, " ")

                                fmt.Println("Arguments: ", cmd_args_split)

                                syscallexec.ExecSysCall(cmd_str, cmd_args_split...)
                                // syscallexec.ExecProcess(cmd_str, verbose, cmd_args_split...)
                            } else {
                                syscallexec.ExecSysCall(cmd_str)
                                // syscallexec.ExecProcess(cmd_str, verbose)
                            }
                        }
                    case "execute-system-command-realtime":
                        fmt.Println("Verbose: ", verbose)

                        // Check if "syscall-set-command" is in 'opt_with_Arguments' and Obtain the command line string to execute
                        if cmd_str, is_found := opt_with_Arguments["syscall-set-command"]; is_found {
                            fmt.Println("Command: ", cmd_str)

                            // Check if 'syscall-set-arguments" is in 'opt_with_Arguments' and obtain the arguments
                            if cmd_args, is_found := opt_with_Arguments["syscall-set-arguments"]; is_found {
                                // Split the cmd_args string into a list
                                var cmd_args_split []string = strings.Split(cmd_args, " ")

                                fmt.Println("Arguments: ", cmd_args_split)

                                syscallexec.ExecSysCallRealtime(cmd_str, verbose, cmd_args_split...)
                            } else {
                                syscallexec.ExecSysCallRealtime(cmd_str, verbose)
                            }
                        }
                    case "print-json":
                        /*
                         * Perform JSON Parsing
                         */
                        var json_fpathname string = "test.json"

                        // Check if a JSON file is found
                        if _, err := os.Stat(json_fpathname); err == nil {
                            // Open JSON file
                            var json_fptr *os.File = jsonio.OpenFile(json_fpathname)

                            // Read the opened JSON file
                            jsonio.ReadJSON(json_fptr)

                            // Get the JSON file contents
                            //var json_fstruct jsonio.JSONStruct = jsonio.GetJSONStruct()
                            var json_fcontents = jsonio.GetJSONContents()

                            // Printing out JSON file contents
                            fmt.Println("")
                            fmt.Println("JSON File Contents:")
                            fmt.Println("")
                            for k,v := range json_fcontents {
                                fmt.Println(k, ":", v)
                            }

                            // Close the opened JSON file after usage
                            jsonio.CloseFile(json_fptr)
                        } else {
                            fmt.Println(err)
                        }
                    case "start-todolist-webserver":
                        // Startup the SQLite3 TODO list webserver
                        sqlite3db.StartWebServer("", -1, "", "", "")
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
}

