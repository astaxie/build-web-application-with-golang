## Set up for `ch.5.3`

- Step 1) Download and install sqlite 3.
- Step 2) Run `sqlite3 foo.db` to create a databased called `foo`.
- Step 3) Create the `userinfo` table in sqlite using `schema.sql`.

	Read and run sql statements 

		sqlite> .read schema.sql

	Show tables

		sqlite> .tables
		userinfo


- Step 4) Exit sqlite.

		sqlite> .exit

- Step 5) Run `go get` to download and install remote packages.
- Step 6) Run the program with `go run main.go`

