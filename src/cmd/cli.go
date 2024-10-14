/*
 * Module containing the core CLI argument parsing and handling logic
 */
package cmd

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

            // Check if there is a split (means '=' exists)
            if len(curr_arg_split) > 1 {
                // There is a split
                curr_arg_key = curr_arg_split[0]
                curr_arg_value = curr_arg_split[1]
            } else {
                curr_arg_key = curr_arg_split[0]
            }

            // Process the argument
            switch curr_arg_key {
                // If the CLI argument has a value, Reduce the index by 1
                case "-h", "--help":
                    res_args_opt_flags["help"] = true
                case "-v", "--version":
                    res_args_opt_flags["version"] = true
                case "-p", "--print-all-arguments":
                    // Print all arguments
                    res_args_opt_flags["print-all-arguments"] = true
                case "-V", "--verbose":
                    // Enable verbose output
                    res_args_opt_flags["verbose"] = true
                case "--set-cmd-str":
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
                                res_args_opt_with_Value["syscall-set-command"] = ""

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
                        res_args_opt_with_Value["syscall-set-command"] = curr_arg_value
                    }
                case "--set-cmd-args":
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
                                res_args_opt_with_Value["syscall-set-arguments"] = ""

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
                        res_args_opt_with_Value["syscall-set-arguments"] = curr_arg_value
                    }
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

