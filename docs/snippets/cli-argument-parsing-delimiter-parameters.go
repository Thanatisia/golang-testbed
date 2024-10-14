/*
 * Module containing the core CLI argument parsing and handling logic
 */
package cli

/*
 * Import packages/modules/libraries
 */
import (
	"fmt"
	"os"
	"strings"
	// "maps" // Built-in Standard Library module for handling Key-Value Mappings (aka Dictionary, Associative Array, HashMap, Maps etc etc)
)

/*
 * Define functions to export
 */
func ParseArguments() (map[string]bool, map[string]string, []string, []string) {
    /*
     * Parse CLI arguments from scratch
     */

    // Initialize Variables
    var exec string = os.Args[0]
    var argv []string = os.Args[1:]
    var argc int = len(argv)
    var res_args_opt_flags map[string]bool = make(map[string]bool)
    var res_args_opt_with_Value map[string]string = make(map[string]string)
    var res_args_pos []string
    var error_msg []string

    fmt.Println("Executable: ", exec)

    // Check if there are CLI arguments provided
    if (argc > 0) {
        i := 0
        // Iterate through all the parsed CLI arguments until there are no arguments remaining and process them; For is golang's 'while loop'
        for i < argc {
            // Get the current argument
            var curr_arg_tmp string = argv[i];

            // Initialize Variables
            var delimiter string = "="
            var curr_arg_key string
            var curr_arg_value string

            // Check if there is a '=' delimiter (Split the string to 'variable-key-name, delimiter, variable-value')
            var curr_arg_split []string = strings.Split(curr_arg_tmp, delimiter)
            fmt.Println(curr_arg_split)

            // Check if there is a split (means '=' exists)
            if len(curr_arg_split) > 1 {
                // There is a split
                curr_arg_key = curr_arg_split[0]
                curr_arg_value = curr_arg_split[1]
            } else {
                curr_arg_key = curr_arg_split[0]
            }

            fmt.Println("Delimiter: ", delimiter)
            fmt.Println("Current Argument: ", curr_arg_key)
            fmt.Println("Argument Value: ", curr_arg_value)

            // Process the argument
            switch curr_arg_key {
                // If the CLI argument has a value, Reduce the index by 1
                case "-h", "--help":
                    res_args_opt_flags["help"] = true
                case "-v", "--version":
                    res_args_opt_flags["version"] = true
                case "--set-value":
                    // Check if the split contains a value
                    if curr_arg_value == "" {
                        // The split had no 'LHS=RHS' value
                        // Check if arguments are provided
                        if i < argc-1 {
                            // Arguments are provided
                            var next_i int = i+1
                            tmp_val := argv[next_i]

                            // Data Validation: Null Value and Valid Value Check
                            if tmp_val != "" {
                                res_args_opt_with_Value["set-value"] = tmp_val

                                // Increment the index counter by 1 to jump to the next argument
                                i += 1
                            } else {
                                error_msg = append(error_msg, fmt.Sprintf("Parameter [%s]: Value not provided.", curr_arg_key))
                            }
                        } else {
                            error_msg = append(error_msg, fmt.Sprintf("Parameter [%s]: Insufficient values provided.", curr_arg_key))
                        }
                    } else {
                        // Split had a 'LHS=RHS' value
                        res_args_opt_with_Value["set-value"] = curr_arg_value
                    }
                default:
                    res_args_pos = append(res_args_pos, curr_arg_key)
            }

            // Increment the index by 1 to go to the next argument
            i += 1;
        }
    } else {
        error_msg = append(error_msg, "No arguments provided.")
    }

    // Return/Output
    return res_args_opt_flags, res_args_opt_with_Value, res_args_pos, error_msg
}

// Define main() function
func main() {
    // Parse CLI arguments from the user into the system
    opts_Flags, opt_with_Arguments, positionals, error_msg := ParseArguments()

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

                // Process the key-value mappings
                switch key {
                    case "set-value":
                        fmt.Println("Value Received: ", value)
                    default:
                        fmt.Printf("Invalid optional argument provided: %s\n", key)
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
}

