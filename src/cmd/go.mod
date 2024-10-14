module testbench_practiceground/cmd

go 1.21.5

replace testbench_practiceground/sqlite3db => ../modules/features/sqlite3db

replace testbench_practiceground/hello => ../modules/tutorials/hello

require testbench_practiceground/hello v0.0.0-00010101000000-000000000000
