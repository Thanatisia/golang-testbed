/*
 * SQLite3 Database Connection Test Module
 */
package sqlite3db

import (
    "database/sql"  // Built-in system library/package for SQL relational database interactions
    // "errors"     // Built-in system library/package for error checking
    // "log"        // Built-in system library/package for logging
    "sync"
    // "fmt"        // Built-in system library/package for formatted I/O

    // External Dependencies
	_ "github.com/mattn/go-sqlite3" // SQLite driver
)

/*
 * Type definitions (typedef) - Structure/Struct container
 */
type DBTable struct {
    mu sync.Mutex; // Mutual Exclusive Lock
    db *sql.DB;    // SQL Database connection pointer object
}

// NOTE: Functions in golang need to start with a capital letter, otherwise it will not export
func Connect(db_file string) (*DBTable, error) {
    /*
     * Initialize a new Database Connection and return the database table structure object containing the opened database handler point object
     *
     * :: Return
     * - Type: (DBTable structure, error)
     */

    // Data Validation: Null Value Check
    if db_file == "" {
        db_file = "sqlite.db";
    }

    // Open SQLite database connection handler/object
    db, err := sql.Open("sqlite3", db_file);

    // Error validation check
    if err != nil {
        return nil, err;
    }

    /*
     * Initialize a new instance of the DBTable data structure
     * De-reference and retrieve the DBTable struct from the memory address, and populate it with the database handler/pointer object
     */
    new_dbTable := &DBTable {
        // structure-field: new-value
        db: db,
    }

    // Output/Return
    return new_dbTable, nil
}

func ExecSQLStmt(db_conn *DBTable, sql_stmt string) (int64, int64, error) {
    /*
     * Execute a non-query SQL Statement to the database and return the last affected row ID (if any)
     *
     * :: Return
     * - Type: (int64, int64, error)
     * - values: (last_id, row_count, error)
     */

    // Define variables
    var res sql.Result;
    var last_id int64;
    var row_count int64;

    // Retrieve the database object
    db := db_conn.db;

    // Execute SQL statement and return the result
    res, err := db.Exec(sql_stmt);

    // Obtain the last affected row ID
    last_id, err = res.LastInsertId();

    // Error validation
    if err != nil {
        return 0, 0, err
    }

    // Obtain the number of rows returned
    row_count, err = res.RowsAffected();

    // Error validation
    if err != nil {
        return last_id, 0, err
    }

    // Output/Return
    return last_id, row_count, nil
}

func ExecQuerySQLStmt(db_conn *DBTable, sql_stmt string, fields ...any) (string, int64, error) {
    /*
     * Execute a query (retrieval/select) SQL Statement to the database and return the query results (if any)
     *
     * :: Return
     * - Type: (string, int64, error)
     * - values: (rows, row_count, error)
     */

    // Define variables
    var row_cursor *sql.Row;
    var rows_cursor *sql.Rows;
    var rows string;
    var row_count int64;

    // Retrieve the database object
    db := db_conn.db;

    // Execute SQL statement and return the resulting rows from the query
    rows_cursor, err := db.Query(sql_stmt);

    // Error validation
    if err != nil {
        return "", 0, err
    }

    // Obtain the queried rows returned
    row_cursor = db.QueryRow(sql_stmt);

    // Error validation
    if err != nil {
        return "", 0, err
    }

    // Obtain the number of rows returned
    var err error = row_cursor.Scan(&fields);

    // Error validation
    if err != nil {
        return sql_Row, 0, err
    }

    // Output/Return
    return last_id, row_count, nil
}

func RetrieveQuery(db_conn *DBTable, row *sql.Row) (DBTable, error) {
    /*
     * Receive the rows from the database cursor
     *
     * :: Return
     * - Type: (string, int64, error)
     * - values: (rows, row_count, error)
     */

    // Scan row
    err := row.Scan(&db_conn.db.Description);
}

func Disconnect(db_conn *DBTable) (*DBTable, error) {
    /*
     * Close the specified Database Connection of the database connection pointer handler
     *
     * :: Return
     * - Type: (DBTable structure, error)
     */

    // Declare variables
    var err error;

    // Data Validation: Null Value Check
    if db_conn.db != nil {
        // Not null, close database
        err = db_conn.db.Close();
    }

    // Error validation check
    if err != nil {
        return db_conn, err;
    } else {
        return nil, err;
    }
}

