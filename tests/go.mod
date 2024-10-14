module tests

go 1.21.5

replace tests/cli => ./src/cli

replace tests/cmd => ./src/cmd

require (
	tests/cli v0.0.0-00010101000000-000000000000
	tests/cmd v0.0.0-00010101000000-000000000000
)
