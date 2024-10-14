module golang_testbed/cmd

go 1.21.5

replace golang_testbed/sqlite3db => ../modules/features/sqlite3db

replace golang_testbed/hello => ../modules/tutorials/hello

require golang_testbed/hello v0.0.0-00010101000000-000000000000
