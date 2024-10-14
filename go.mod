module golang_testbed

go 1.21.5

replace golang_testbed/cmd => ./src/cmd

replace golang_testbed/sqlite3db => ./src/modules/features/sqlite3db

replace golang_testbed/system_cmd_execution => ./src/modules/features/system_cmd_execution

replace golang_testbed/hello => ./src/modules/tutorials/hello

require (
	golang_testbed/cmd v0.0.0-00010101000000-000000000000
	golang_testbed/system_cmd_execution v0.0.0-00010101000000-000000000000
)
