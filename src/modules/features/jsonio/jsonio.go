/*
 * JSON I/O Library
 */
package jsonio

// Import dependencies
import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// Define a new type
type JSONStruct struct {
    // Attributes|Property|Variable <data-type>
    json_fname string;  // JSON file name; Type: String
    json_fptr *os.File; // JSON File Pointer/Header; Type: *os.File (os.File pointer)
    json_fcontents map[string]any; // JSON File Contents; Type Map<String,any>
}

/*
 * Define global variables/structures
 */
// Initialize a new struct object instance
var json_fstruct JSONStruct

/*
 * Declare functions
 */

// Getter functions
func GetJSONStruct() (JSONStruct) { return json_fstruct }
func GetJSONFilePtr() (*os.File) { return json_fstruct.json_fptr }
func GetJSONName() (string) { return json_fstruct.json_fname }
func GetJSONContents() (map[string]any) { return json_fstruct.json_fcontents }

func OpenFile(json_filename string) (*os.File) {
    /*
     * Open JSON File and return the json file object pointer/header
     */

    // Open JSON file
    f_json, err := os.Open(json_filename)

    // Data Validation: Error Checking
    if err != nil {
        fmt.Println(err)
    }

    // Add the file name to the structure
    json_fstruct.json_fname = json_filename

    // Add the file pointer to the structure
    json_fstruct.json_fptr = f_json

    // Return the file pointer
    return f_json
}

func ReadJSON(f_json *os.File) {
    /*
     Read an opened JSON file pointer and unmarshal the JSON object
     */

    // Read the opened file pointer as a byte array
    byte_val, err := io.ReadAll(f_json)

    // Perform error checking
    if err != nil {
        fmt.Println(err)
    }

    // Unmarshal (Extract/Uncompress) the byte array (which contains the JSON file contents) into the value in struct object pointer's memory address
    json.Unmarshal(byte_val, &json_fstruct.json_fcontents)

    // Store the Byte array in the struct
    // json_fstruct.json_fcontents = byte_val
}

func CloseFile(f_json *os.File) {

    // Check if the file pointer is empty
    if f_json != nil {
        // Defer the closing of our JSON file so that we can parse it later on
        defer f_json.Close()

        // Set the file pointer to empty
        json_fstruct.json_fname = ""
        json_fstruct.json_fptr = nil
        json_fstruct.json_fcontents = nil
    } else {
        fmt.Println("JSON file is not opened")
    }
}


