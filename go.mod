module testbench_practiceground

go 1.21.5

replace testbench_practiceground/cmd => ./src/cmd

replace testbench_practiceground/sqlite3db => ./src/modules/features/sqlite3db

replace testbench_practiceground/system_cmd_execution => ./src/modules/features/system_cmd_execution

replace testbench_practiceground/hello => ./src/modules/tutorials/hello

require (
	testbench_practiceground/cmd v0.0.0-00010101000000-000000000000
	testbench_practiceground/system_cmd_execution v0.0.0-00010101000000-000000000000
)
