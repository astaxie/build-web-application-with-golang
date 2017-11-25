## Setup for ch.5.4

- Step 1) Install and run Postgres
- Step 2) Create a user and database according to the constants in `main.go`
	
		DB_USER     = "user"
		DB_PASSWORD = ""
		DB_NAME     = "test"

- Step 3) Create table `userinfo` located at `schema.sql`
- Step 4) Run `go get` to download and install the remote packages.
- Step 5) Execute the program with `go run main.go`
