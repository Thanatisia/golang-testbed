package sqlite3db

import (
    "fmt"
    // "os"
)

func ApplicationLogic() {
    // Initialize a variable with a specific data type
    // var sql_conn db.DBTable;
    // var err error;
    var db_file string;
    var sql_stmt string;

    // Define database file
    db_file = "sqlite.db"

    // Define SQL statement
    sql_stmt = `
    CREATE TABLE IF NOT EXISTS new_table (
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        description TEXT
    );
    `

    // Connect to database
    sql_conn, err := Connect(db_file);

    if err != nil {
        // Error connecting
        fmt.Println("Error connecting to database: ", err);
    } else {
        // Connected successfully
        fmt.Println("Database connected successfully: ", db_file);

        // Create table if it doesnt exist
        last_id, row_count, err := ExecSQLStmt(sql_conn, sql_stmt);

        if err != nil {
            fmt.Println("Rows affected: ", row_count);
            fmt.Println("Last affected ID: ", last_id);
        } else {
            fmt.Println("Error: ", err);
        }

        // Select all values
        sql_stmt = `
        SELECT * FROM new_table;
        `

        // Execute SQL statements
        last_id, row_count, err := ExecQuerySQLStmt(sql_conn, sql_stmt)

        // Close connection
        sql_conn, err = Disconnect(sql_conn)

        // Check if disconnection was successful
        if sql_conn == nil {
            fmt.Println("Database disconnected successfully");
        } else {
            fmt.Println("Error disconnecting database: ", err);
        }
    }
}

