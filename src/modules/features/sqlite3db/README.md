# Golang Cookbook/Snippet - TODO List with SQLite3 + HTTP Server + Web Templates

## Setup

### Dependencies
+ go

### Pre-Requisites
- Set Environment Variables
    - `GOROOT` : Set this to the root directory of your golang SDK
        ```bash
        export GOROOT=/usr/local/go
        ```
    - `GOPATH` : Set this to the root directory of your golang home configuration directory
        ```bash
        export GOPATH=$HOME/.go
        ```
    - `GOBIN`  : Set this to the target directory to install/save the installed binaries into; Default: $HOME/.go/bin
        ```bash
        export GOBIN=$GOPATH/bin
        ```

- Setup project package workspace directory
    - Make package root directory
        ```bash
        mkdir -pv path/to/project-name
        ```
    - Change directory to package root directory
        ```bash
        cd /path/to/project-name
        ```
    - Initialize the go package module
        ```bash
        go mod init [package-name]
        ```
    - Create the main.go entry point source file
        ```bash
        cat <<EOF > main.go
        package main

        import (
            "database/sql"  // Built-in system library/package for SQL relational database interactions
            "fmt"           // Built-in system library/package for formatted I/O
            "log"           // Built-in system library/package for logging
            "net/http"      // Built-in system library/package for HTTP client and server
            "text/template" // Built-in system library/package for HTML templates

            // External Dependencies
            _ "github.com/mattn/go-sqlite3" // SQLite driver
        )

        func main() {
            fmt.Println("Hello World")
        }
        EOF
        ```
    - Download/Pull/"Get" the golang SQLite driver library
        ```bash
        go get github.com/mattn/go-sqlite3
        ```
    - Update/"tidy" the go package module definition file
        ```bash
        go mod tidy
        ```

## Documentations

### Components
+ SQLite3 Database I/O Processing and CRUD webserver component functions
+ HTML Template and Rendering
+ Serve Webserver (Startup the Webserver and render the HTML Template)
+ HTTP Server Routing (Serve and Listen for HTTP Request and Response)

### Snippet Explanation
- Import package/library dependencies
    ```go
    import (
        "database/sql"  // Built-in system library/package for SQL relational database interactions
        "fmt"           // Built-in system library/package for formatted I/O
        "log"           // Built-in system library/package for logging
        "net/http"      // Built-in system library/package for HTTP client and server
        "text/template" // Built-in system library/package for HTML templates

        // External Dependencies
        _ "github.com/mattn/go-sqlite3" // SQLite driver
    )
    ```

- Type definitions (typedef) - Structure/Struct container
    - Explanation
        - `type Todo struct { ... }` : Define a new struct object and declare it as a data type alias name 'Todo'
            + This is similar to 'typedef' in C which will effectively map the specified new type as an alias to the struct object
            - `variable_name <datatype>;` : Specify an attribute/property/variable in the struct definition/container
    ```go
    type Todo struct {
        ID int; // To store the row ID of the list entry
        Title string; // To store the title of the list entry
    }
    ```

- Initialize global variables
    - SQLite database connection handler/pointer object of type '*sql.DB'
        ```go
        var DB *sql.DB;
        ```

- Define private functions
    - `initDB(db_type string, db_fname string)` : Open and initialize Database connection and return the database connection object
        - Params
            - db_type : Specify the database connection type to open
                + Type: String
                + Default: sqlite3
            - db_fname : Specify the database file name to open
                + Type: String
                + Default: app.db
        ```go
        func initDB(db_type string, db_fname string) { ... }
        ```
        - Initialize local variables
            ```go
            var err error; // Error handler
            var sql_stmt = `
            CREATE TABLE IF NOT EXISTS todos (
                id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
                title TEXT
            );
            `
            ```
        - Data Validation: Null Value Check
            ```go
            if db_type == "" {
                db_type = "sqlite3"
            }
            if db_fname == "" {
                db_fname = "app.db"
            }
            ```
        - Open a connection to the SQLite database file and return the database connection handler and the error status
            ```go
            DB, err = sql.Open(db_type, db_fname)
            if err != nil {
                // Log an error and stop the program if the database cant be opened properly
                log.Fatal(err)
            }
            ```
        - Execute the SQLite statement and create a database table if it doesnt exist
            ```go
            _, err = DB.Exec(sql_stmt) // This will execute the SQL query statement without returning any rows
            if err != nil {
                // Log and error if the table cannot be created
                log.Fatalf("Error creating table: %q: %s\n", err, sql_stmt); // log.Fatalf is equivalent to Printf for logging
            }
            ```
    - `indexHandler(http_res_writer http.ResponseWriter, http_req *http.Request)` : Handles the creation of an index webpage and Serves a Webserver and displays the main page with all the TODO list entries
        - Params
            - http_res_writer : Pass the HTTP Response Writer object
                - Type: http.ResponseWriter
            - http_req : HTTP Request Handler pointer object
                - Type: *http.Request
        ```go
        func indexHandler(http_res_writer http.ResponseWriter, http_req *http.Request) { ... }
        ```
        - Initialize Variables
            ```go
            var sql_stmt string = "SELECT id, title FROM todos"
            ```
        - Query the database to get all TODO list entries
            ```go
            rows, err := DB.Query(sql_stmt)
            if err != nil {
                // Throw an HTTP Server Error status and
                // Print the HTTP Response Writer object, the error message and
                // Return an HTTP 500 error status code if the query fails
                http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
                return
            }
            ```
        - Ensure that the rows are closed after processing before proceeding
            ```go
            defer rows.Close()
            ```
        - Process the TODO list retrieved and store them into a todos object
            ```go
            todo_list := []Todo{} // Slice to store TODO list entries
            ```
        - Scan all rows returned from the query and prepare the Next result row for reading
            ```go
            for rows.Next() {
            ```
        - Initialize a new TODO list entry of type structure 'Todo'
            ```go
                var curr_todolist_entry Todo
            ```
        - Scan the current row and attempt to find & obtain the specified columns, and
            - If columns are found: store the query results obtained corresponding to the column into the current Todo list instance
                ```go
                if err := rows.Scan(&curr_todolist_entry.ID, &curr_todolist_entry.Title); err != nil {
                    // Throw an HTTP Server Error status and
                    // Print the HTTP Response Writer object, the error message and
                    // Return an HTTP 500 error status code if the query fails
                    http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
                    return
                }
                ```
            - Append the current TODO list entry into the master list
                ```go
                todo_list = append(todo_list, curr_todolist_entry)
                ```
        - End the iterator block
            ```go
            }
            ```

        - Prepare and Initialize a new HTML Template and Parse in the HTML file content string into the HTML Template object
             - Notes
                - Use '{{.<variable>}}' to use backend data/variables in the frontend DOM
                    - This is similar to
                        + how Javascript frameworks like React and Svelte interact with the HTML webapp components
                        + Python flask - Using python variables in the frontend HTML file
                - The id of each tag can be obtained in Golang by retrieving it using the HTTP Request object (http_req.TagValue("your-id"))
            ```go
            tmpl := template.Must(template.New("index").Parse(`
               <!DOCTYPE html>
               <html>
                   <head>
                       <title>Todo List</title>
                   </head>
                   <body>
                       <h1>Todo List</h1>
                       <form action="/create" method="POST">
                           <input type="text" name="title" placeholder="New Todo" required>
                           <button type="submit">Add</button>
                       </form>
                       <ul>
                           {{range .}}
                           <li>{{.Title}} <a href="/delete?id={{.ID}}">Delete</a></li>
                           {{end}}
                       </ul>
                   </body>
               </html>
            `))
            ```

        - Parse and execute the HTML template with the TODO list dataset and Render the template with the list of TODO list entries
            ```go
            tmpl.Execute(http_res_writer, todo_list)
            ```
    - `createHandler(http_res_writer http.ResponseWriter, http_req *http.Request)` : Handles the creation of a new TODO list entry (CREATE)
        - Params
            - http_res_writer : Pass the HTTP Response Writer object
                - Type: http.ResponseWriter
            - http_req : HTTP Request Handler pointer object
                - Type: *http.Request
        ```go
        func createHandler(http_res_writer http.ResponseWriter, http_req *http.Request) { ... }
        ```
        - Receive the HTTP Request and check the request method
            ```go
            var request_method string = http_req.Method
            ```

        - Data Validation: Check if the HTTP Request method is a POST request
            ```go
            if request_method == "POST" {
            ```
            - Obtain the title from the title <form> tag in the frontend by referencing the DOM tag ID
                ```go
                title := http_req.FormValue("title")
                ```

            - Format SQL statement
                ```go
                var sql_stmt string = fmt.Sprintf("INSERT INTO %s(%s) VALUES(?)", "todos", "title")
                ```

            - Insert the form value into the SQLite database table in the form of a prepared statement
                ```go
                _, err := DB.Exec(sql_stmt, title)
                if err != nil {
                    // Throw an HTTP Server Error status and
                    // Print the HTTP Response Writer object, the error message and
                    // Return an HTTP 500 error status code if the query fails
                    http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
                    return
                }
                fmt.Println("Entry [", title, "] has been inserted successfully.")
                ```

            - Redirect to the index (main) page (aka root directory) after the successful creation/insert
                + The main page TODO list should now display the new list entry
                ```go
                http.Redirect(http_res_writer, http_req, "/", http.StatusSeeOther)
                ```
    - `deleteHandler(http_res_writer http.ResponseWriter, http_req *http.Request)` : Handles the deletion of a new TODO list entry (DELETE)
        - Params
            - http_res_writer : Pass the HTTP Response Writer object
                - Type: http.ResponseWriter
            - http_req : HTTP Request Handler pointer object
                - Type: *http.Request
        ```go
        func deleteHandler(http_res_writer http.ResponseWriter, http_req *http.Request) { ... }
        ```
        - Get the id from the first value of the URL query parameters of the HTTP Request
            ```
            entry_id := http_req.URL.Query().Get("id")
            ```

        - Format SQL statement
            ```go
            var sql_stmt string = fmt.Sprintf("DELETE FROM %s WHERE id = ?", "todos")
            ```

        - Delete the selected TODO list entry from the database table
            ```go
            _, err := DB.Exec(sql_stmt, entry_id)
            if err != nil {
                // Throw an HTTP Server Error status and
                // Print the HTTP Response Writer object, the error message and
                // Return an HTTP 500 error status code if the query fails
                http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
                return
            }
            fmt.Println("Entry [", entry_id, "] has been deleted successfully.")
            ```

        - Redirect to the index (main) page (aka root directory) after the successful creation/insert
            + The main page TODO list should now display the new list entry
            ```go
            http.Redirect(http_res_writer, http_req, "/", http.StatusSeeOther)
            ```

- Define public functions
    - `StartWebServer(http_server_ip_address string, http_server_port_number int, http_server_protocol string, db_fname string, db_type string)` : Startup the HTTP Webserver
        ```go
        func StartWebServer(http_server_ip_address string, http_server_port_number int, http_server_protocol string, db_fname string, db_type string) { ... }
        ```
        - Data Validation: Null Value Check
            ```go
            if http_server_ip_address == "" {
                http_server_ip_address = "127.0.0.1"
            }
            if http_server_port_number < 0 {
                http_server_port_number = 8080
            }
            if http_server_protocol == "" {
                http_server_protocol = "http"
            }
            ```

        - Initialize Variable
            ```go
            // var server_addr_url string = fmt.Sprintf("%s://%s:%d", http_server_protocol, http_server_ip_address, http_server_port_number) // http_server_protocol + "://" + http_server_ip_address + ":" + string(http_server_port_number)
            var server_addr_url string = fmt.Sprintf("%s:%d", http_server_ip_address, http_server_port_number) // http_server_protocol + "://" + http_server_ip_address + ":" + string(http_server_port_number)
            ```

        - Initialize the Database
            ```go
            initDB(db_type, db_fname)
            ```

        - Ensure that the database connection is closed when the program exits
            - Notes
                - defer: Ensure that the command specified is performed when the program exits
            ```go
            defer DB.Close()
            ```

        - Route the Handlers for each URL path
            - Explanation
                - This is equivalent to HTTP Webserver (Location) Routing
                - You are setting your HTTP REST API webserver routes and the callback function to execute when the event handler is triggered
            ```go
            http.HandleFunc("/", indexHandler) // When 'http://[server-ip-address]:[server-port-number]/' (index page/root directory) is accessed
            http.HandleFunc("/create", createHandler) // When 'http://[server-ip-address]:[server-port-number]/create' is accessed
            http.HandleFunc("/delete", deleteHandler) // When 'http://[server-ip-address]:[server-port-number]/delete' is accessed
            ```

        - Startup the HTTP Webserver using the specified (IP address and Port number) socket and start routing
            ```go
            fmt.Println("Server is running at", server_addr_url)
            log.Fatal(http.ListenAndServe(server_addr_url, nil))
            ```

## Wiki

### Code

```go
/*
 * TODO List implementation using sqlite3db as a core driver library
 * :: References
 * - https://medium.com/@peymaan.abedinpour/golang-crud-app-tutorial-step-by-step-guide-using-sqlite-a3ce08a4fc81
 */
package main

// Import Libraries/Package Dependencies
import (
    "database/sql"  // Built-in system library/package for SQL relational database interactions
    "fmt"           // Built-in system library/package for formatted I/O
    "log"           // Built-in system library/package for logging
    "net/http"      // Built-in system library/package for HTTP client and server
    "text/template" // Built-in system library/package for HTML templates

    // External Dependencies
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

/*
 * Type definitions (typedef) - Structure/Struct container
 */
type Todo struct {
    ID int; // To store the row ID of the list entry
    Title string; // To store the title of the list entry
}

/* Initialize global variables */

// Initialize SQLite database connection handler/pointer object
var DB *sql.DB;

/*
 * Define private functions
 */
func initDB(db_type string, db_fname string) {
    /*
     * Open and initialize Database connection and return the database connection object
     * :: Params
     * - db_type : Specify the database connection type to open
     *      + Type: String
     *      + Default: sqlite3
     * - db_fname : Specify the database file name to open
     *      + Type: String
     *      + Default: app.db
     */

    // Initialize local variables
    var err error; // Error handler
    var sql_stmt = `
    CREATE TABLE IF NOT EXISTS todos (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        title TEXT
    );
    `

    // Data Validation: Null Value Check
    if db_type == "" {
        db_type = "sqlite3"
    }
    if db_fname == "" {
        db_fname = "app.db"
    }

    // Open a connection to the SQLite database file and return the database connection handler and the error status
    DB, err = sql.Open(db_type, db_fname)
    if err != nil {
        // Log an error and stop the program if the database cant be opened properly
        log.Fatal(err)
    }

    // Execute the SQLite statement and create a database table if it doesnt exist
    _, err = DB.Exec(sql_stmt) // This will execute the SQL query statement without returning any rows
    if err != nil {
        // Log and error if the table cannot be created
        log.Fatalf("Error creating table: %q: %s\n", err, sql_stmt); // log.Fatalf is equivalent to Printf for logging
    }
}

func indexHandler(http_res_writer http.ResponseWriter, http_req *http.Request) {
    /*
     * Serves a Webserver and displays the main page with all the TODO list entries
     * :: Params
     * - http_res_writer : Pass the HTTP Response Writer object
     *      - Type: http.ResponseWriter
     * - http_req : HTTP Request Handler pointer object
     *      - Type: *http.Request
     */
     // Initialize Variables
     var sql_stmt string = "SELECT id, title FROM todos"

     // Query the database to get all TODO list entries
     rows, err := DB.Query(sql_stmt)
     if err != nil {
         // Throw an HTTP Server Error status and
         // Print the HTTP Response Writer object, the error message and
         // Return an HTTP 500 error status code if the query fails
         http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
         return
     }
     // Ensure that the rows are closed after processing before proceeding
     defer rows.Close()

     // Process the TODO list retrieved and store them into a todos object
     todo_list := []Todo{} // Slice to store TODO list entries

     // Scan all rows returned from the query and prepare the Next result row for reading
     for rows.Next() {
        // Initialize a new TODO list entry of type structure 'Todo'
        var curr_todolist_entry Todo

        // Scan the current row and attempt to find & obtain the specified columns, and
        // If columns are found: store the query results obtained corresponding to the column into the current Todo list instance
        if err := rows.Scan(&curr_todolist_entry.ID, &curr_todolist_entry.Title); err != nil {
            // Throw an HTTP Server Error status and
            // Print the HTTP Response Writer object, the error message and
            // Return an HTTP 500 error status code if the query fails
            http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
            return
        }

        // Append the current TODO list entry into the master list
        todo_list = append(todo_list, curr_todolist_entry)
     }

     /*
      * Prepare and Initialize a new HTML Template and Parse in the HTML file content string into the HTML Template object
      * :: Notes
      * - Use '{{.<variable>}}' to use backend data/variables in the frontend DOM
      *     - This is similar to
      *         + how Javascript frameworks like React and Svelte interact with the HTML webapp components
      *         + Python flask - Using python variables in the frontend HTML file
      * - The id of each tag can be obtained in Golang by retrieving it using the HTTP Request object (http_req.TagValue("your-id"))
      */
     tmpl := template.Must(template.New("index").Parse(`
        <!DOCTYPE html>
        <html>
            <head>
                <title>Todo List</title>
            </head>
            <body>
                <h1>Todo List</h1>
                <form action="/create" method="POST">
                    <input type="text" name="title" placeholder="New Todo" required>
                    <button type="submit">Add</button>
                </form>
                <ul>
                    {{range .}}
                    <li>{{.Title}} <a href="/delete?id={{.ID}}">Delete</a></li>
                    {{end}}
                </ul>
            </body>
        </html>
     `))

     // Parse and execute the HTML template with the TODO list dataset
     // and Render the template with the list of TODO list entries
     tmpl.Execute(http_res_writer, todo_list)
}

func createHandler(http_res_writer http.ResponseWriter, http_req *http.Request) {
    /*
     * Handles the creation of a new TODO list entry (CREATE)
     * :: Params
     * - http_res_writer : Pass the HTTP Response Writer object
     *      - Type: http.ResponseWriter
     * - http_req : HTTP Request Handler pointer object
     *      - Type: *http.Request
     */

    // Receive the HTTP Request and check the request method
    var request_method string = http_req.Method

    // Data Validation: Check if the HTTP Request method is a POST request
    if request_method == "POST" {
        // Obtain the title from the title <form> tag in the frontend by referencing the DOM tag ID
        title := http_req.FormValue("title")

        // Format SQL statement
        var sql_stmt string = fmt.Sprintf("INSERT INTO %s(%s) VALUES(?)", "todos", "title")

        // Insert the form value into the SQLite database table in the form of a prepared statement
        _, err := DB.Exec(sql_stmt, title)
        if err != nil {
            // Throw an HTTP Server Error status and
            // Print the HTTP Response Writer object, the error message and
            // Return an HTTP 500 error status code if the query fails
            http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
            return
        }
        fmt.Println("Entry [", title, "] has been inserted successfully.")

        // Redirect to the index (main) page (aka root directory) after the successful creation/insert
        // The main page TODO list should now display the new list entry
        http.Redirect(http_res_writer, http_req, "/", http.StatusSeeOther)
    }
}

func deleteHandler(http_res_writer http.ResponseWriter, http_req *http.Request) {
    /*
     * Handles the deletion of a new TODO list entry (DELETE)
     * :: Params
     * - http_res_writer : Pass the HTTP Response Writer object
     *      - Type: http.ResponseWriter
     * - http_req : HTTP Request Handler pointer object
     *      - Type: *http.Request
     */
     // Get the id from the first value of the URL query parameters of the HTTP Request
     entry_id := http_req.URL.Query().Get("id")

    // Format SQL statement
    var sql_stmt string = fmt.Sprintf("DELETE FROM %s WHERE id = ?", "todos")

    // Delete the selected TODO list entry from the database table
    _, err := DB.Exec(sql_stmt, entry_id)
    if err != nil {
        // Throw an HTTP Server Error status and
        // Print the HTTP Response Writer object, the error message and
        // Return an HTTP 500 error status code if the query fails
        http.Error(http_res_writer, err.Error(), http.StatusInternalServerError)
        return
    }
    fmt.Println("Entry [", entry_id, "] has been deleted successfully.")

    // Redirect to the index (main) page (aka root directory) after the successful creation/insert
    // The main page TODO list should now display the new list entry
    http.Redirect(http_res_writer, http_req, "/", http.StatusSeeOther)
 }

/*
 * Define public functions
 */
func StartWebServer(http_server_ip_address string, http_server_port_number int, http_server_protocol string, db_fname string, db_type string) {
    /*
     * Startup the HTTP Webserver
     */

    // Data Validation: Null Value Check
    if http_server_ip_address == "" {
        http_server_ip_address = "127.0.0.1"
    }
    if http_server_port_number < 0 {
        http_server_port_number = 8080
    }
    if http_server_protocol == "" {
        http_server_protocol = "http"
    }

    // Initialize Variable
    // var server_addr_url string = fmt.Sprintf("%s://%s:%d", http_server_protocol, http_server_ip_address, http_server_port_number) // http_server_protocol + "://" + http_server_ip_address + ":" + string(http_server_port_number)
    var server_addr_url string = fmt.Sprintf("%s:%d", http_server_ip_address, http_server_port_number) // http_server_protocol + "://" + http_server_ip_address + ":" + string(http_server_port_number)

    // Initialize the Database
    initDB(db_type, db_fname)

    // Ensure that the database connection is closed when the program exits
    // defer: Ensure that the command specified is performed when the program exits
    defer DB.Close()

    // Route the Handlers for each URL path
    // - This is equivalent to HTTP Webserver (Location) Routing
    // - You are setting your HTTP REST API webserver routes and the callback function to execute when the event handler is triggered
    http.HandleFunc("/", indexHandler) // When 'http://[server-ip-address]:[server-port-number]/' (index page/root directory) is accessed
    http.HandleFunc("/create", createHandler) // When 'http://[server-ip-address]:[server-port-number]/create' is accessed
    http.HandleFunc("/delete", deleteHandler) // When 'http://[server-ip-address]:[server-port-number]/delete' is accessed

    // Startup the HTTP Webserver using the specified (IP address and Port number) socket and start routing
    fmt.Println("Server is running at", server_addr_url)
    log.Fatal(http.ListenAndServe(server_addr_url, nil))
}

func main() {
    // Startup the Webserver and begin routing the HTTP requests
    StartWebServer("localhost", 8080, "http", "app.db", "sqlite3")
}
```

## Resources

## References
+ [medium - @peymaan.abedinpour - Golang CRUD application using SQLite tutorial Step-by-Step Guide](https://medium.com/@peymaan.abedinpour/golang-crud-app-tutorial-step-by-step-guide-using-sqlite-a3ce08a4fc81)

## Remarks

