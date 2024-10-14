// Define module's package
package main

// Import dependencies and packages here
import (
	"flag" // Standard Library for basic command line argument parsing
	"fmt"  // Standard Library for String Formatting
	"os"   // Standard Library for Operating System platform utilities similar to 'os' in python
    cli "tests/cli"
    // cmd "tests/cmd"
)

// Functions
func parse_args_Flag() {
    // Parse CLI arguments from command line
    wordPtr := flag.String("word", "foo", "Enter a string");
    intPtr := flag.Int("short-integer", 42, "Enter a 32-bit Integer");
    int64Ptr := flag.Int64("long-integer", 420, "Enter a 64-bit Integer");
    boolPtr := flag.Bool("fork", false, "Enter a boolean/fork value (true/false)");

    // Initialize a new string variable
    var str_var string;

    // Obtain/Parse a string variable argument from the CLI argument and store it directly into the value within the memory address of the string variable pointer
    flag.StringVar(&str_var, "str_var", "bar", "Enter a string variable");

    // Begin Parsing CLI arguments during startup and boot-time, starting from 'os.Args[1:]'
    // Basically, this is similar to using 'argv := os.Args[1:]'
    flag.Parse()

    // Print out all argument values
    fmt.Println("Word obtained: ", *wordPtr);
    fmt.Println("32-bit integer obtained:  ", *intPtr);
    fmt.Println("64-bit integer obtained:  ", *int64Ptr);
    fmt.Println("Boolean value  obtained:  ", *boolPtr);
    fmt.Println("String variable obtained: ", str_var);
    fmt.Println("Non-flag command line arguments (Positionals): ", flag.Args());
}

func parse_args_Manual() (string, []string, int) {
    // Initialize argument list and obtain all arguments provided by the user to the CLI parser
    var exec string = os.Args[0]
    var argv []string = os.Args[1:]
    var argc int = len(argv)

    return exec, argv, argc
}

// Define main entry point function
func main() {
    exec, argv, argc := cli.ParseArgsManual();

    // Print environment/system information
    fmt.Println("Executable: ", exec);

    // Check if there are arguments provided
    if (argc > 0) {
        fmt.Println("Command Line Arguments: ")
        // Iterate through all arguments
        for i:=0; i < argc; i++ {
            // Get current argument
            curr_arg := argv[i]

            // Print out current argument
            fmt.Println(" -", i, ":", curr_arg);
        }
    } else {
        fmt.Println("(-) No CLI arguments provided.")
    }
}


