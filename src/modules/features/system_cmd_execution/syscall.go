/*
 * Golang System Command Execution module
 */
package systemcmdexecution

import (
	"bufio" // Built-in standard library for Buffer (Virtual Memory address container) Input/Output functionality
	"fmt"
	"io"
	"log"
	"os"
	"runtime"
	"strings"

	// "os"
	"os/exec"
	"syscall"
)

func GetPlatform() string {
    /*
     * Obtain the current Go Operating System (GOOS) and return
     */
    var go_operating_system string = runtime.GOOS

    return go_operating_system
}

func ExecSysCall(cmd_str string, cmd_argv ...string) {
    /*
     * Execute System Command/Executables
     */

    // fmt.Println("Command: ", cmd_str)
    // fmt.Println("Arguments: ", cmd_argv)

    // Obtain the path of an executable/binary
    path, err := exec.LookPath(cmd_str)

    // Error validation
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(path)

    // Format the command binary and Attempt to execute the command and return the struct 'Command'
    cmd := exec.Command(cmd_str, cmd_argv...)

    // Execute and Obtain the output of the command
    cmd_out, err := cmd.Output()

    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Command Successfully Executed")

    // Polymorphism: Convert the bytes list into string
    cmd_out_str := string(cmd_out[:])
    cmd_out_str_sanitized := strings.Trim(cmd_out_str, "\n")

    // Split the command line outputs by newline
    cmd_out_split := strings.Split(cmd_out_str_sanitized, "\n")

    fmt.Println("")

    // Iterate through the command output list and print each line out
    for i:=0; i < len(cmd_out_split); i++ {
        // Get current command output byte
        var curr_byte string = cmd_out_split[i]

        // Print output
        fmt.Println(i, ":", curr_byte)
    }
}

func GenerateNewScanner(std_data_stream_pipe io.Reader) (*bufio.Scanner) {
    /*
     * Generate new scanner instance and buffer container to store the data stream(s) obtained/stored by the scanner
     */
    // Initialize a new Scanner instance, Create a new Buffer Memory container for storing the standard output/error data stream piped out from the command, and start scanning for new data and storing into the buffer
    scanner := bufio.NewScanner(std_data_stream_pipe)

    // Scan the buffer (containing the stdout/stderr data stream) for words/string values and set the split function using the obtained list of words as a scanner function
    scanner.Split(bufio.ScanWords)

    // Return the scanner
    return scanner
}

func ExecSysCallRealtime(cmd_str string, enable_print bool, cmd_argv ...string) {
    /*
     * Execute System Command/Executables with real time output
     */

    // Initialize Variables
    var stdout_str_arr []string
    var stderr_str_arr []string

    // fmt.Println("Command: ", cmd_str)
    // fmt.Println("Arguments: ", cmd_argv)

    // Obtain the path of an executable/binary
    path, err := exec.LookPath(cmd_str)

    // Error validation
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(path)

    // Format the command binary and Attempt to execute the command and return the struct 'Command'
    cmd := exec.Command(cmd_str, cmd_argv...)

    // Execute, create a pipe pointing to the standard output (stdout) data stream and redirect the data stream to stdout when the command starts
    stdout, err := cmd.StdoutPipe()

    if err != nil {
        fmt.Println(err)
    }

    // Execute, create a pipe pointing to the standard error (stderr) data stream and redirect the data stream to stderr when the command starts
    stderr, err := cmd.StderrPipe()

    if err != nil {
        fmt.Println(err)
    }

    // Start the command to create a process handler and begin running the command
    cmd.Start()

    fmt.Println("[...] Generating a new Scanner for Standard Output")
    /*
    // Initialize a new Scanner instance, Create a new Buffer Memory container for storing the standard error data stream piped out from the command, and start scanning for new data and storing into the buffer
    scanner_stdout := bufio.NewScanner(stdout)

    // Scan the buffer (containing the stderr data stream) for words/string values and set the split function using the obtained list of words as a scanner function
    scanner_stdout.Split(bufio.ScanWords)
    */
    scanner_stdout := GenerateNewScanner(stdout)

    // Check if a scanner is created
    if scanner_stdout != nil {
        fmt.Println("[+] A new scanner for standard output has been created successfully.")
    } else {
        fmt.Println("[-] Error creating a new scanner for standard output.")
    }

    fmt.Println("[...] Generating a new Scanner for Standard Error")
    /*
    // Initialize a new Scanner instance, Create a new Buffer Memory container for storing the standard error data stream piped out from the command, and start scanning for new data and storing into the buffer
    scanner_stderr := bufio.NewScanner(stderr)

    // Scan the buffer (containing the stderr data stream) for words/string values and set the split function using the obtained list of words as a scanner function
    scanner_stderr.Split(bufio.ScanWords)
    */
    scanner_stderr := GenerateNewScanner(stderr)

    // Check if a scanner is created
    if scanner_stderr != nil {
        fmt.Println("[+] A new scanner for standard error has been created successfully.")
    } else {
        fmt.Println("[-] Error creating a new scanner for standard error.")
    }

    /*
     * Scan for Standard Output data stream
     */

    // Start scanning the scanner and while there are still data within the data stream (the next token is not empty), continue
    fmt.Println("[...] Scanning for Standard Output")
    i := 0
    for scanner_stdout.Scan() {
        // Get the current scanned text
        curr_scanned_Text := scanner_stdout.Text()

        // Check if 'enable_print' is True
        if enable_print == true {
            // Print out the message
            fmt.Println("Standard Output [", i, "]: ", curr_scanned_Text)
        }

        stdout_str_arr = append(stdout_str_arr, fmt.Sprintf("%s ", curr_scanned_Text))

        // Increment counter
        i += 1
    }

    /*
     * Scan for Standard Error data stream
     */

    // Start scanning the scanner and while there are still data within the data stream (the next token is not empty), continue
    fmt.Println("[...] Scanning for Standard Error")
    i = 0
    for scanner_stderr.Scan() {
        // Get the current scanned text
        curr_scanned_Text := scanner_stderr.Text()

        // Check if 'enable_print' is True
        if enable_print == true {
            // Print out the message
            fmt.Println("Standard Error [", i, "]: ", curr_scanned_Text)
        }

        stderr_str_arr = append(stderr_str_arr, fmt.Sprintf("%s ", curr_scanned_Text))

        // Increment counter
        i += 1
    }

    // Wait for the command process, as well as any copying to and from the standard input, output and error data streams to end
    cmd.Wait()

    fmt.Println("Standard Output: ", strings.Join(stdout_str_arr, ""))
    fmt.Println("Standard Error: ", strings.Join(stderr_str_arr, ""))

    fmt.Println("Command Successfully Executed")
}

func ExecProcess(cmd_str string, enable_print bool, cmd_argv ...string) {
    /*
     * Execute system command process using syscall
     */
    bin_path, err := exec.LookPath(cmd_str)
    if err != nil {
        log.Fatal(err)
    }
    env := os.Environ()
    err = syscall.Exec(bin_path, cmd_argv, env)
    if err != nil {
        log.Fatal(err)
    }
}

