**sqlgen** generates SQL statements and database helper functions from your Go structs. It can be used in place of a simple ORM or hand-written SQL. See the [demo](https://github.com/ajiyoshi-vg/sqlgen/tree/master/demo) directory for examples.

### Install

Install or upgrade with this command:

```
go get -u github.com/ajiyoshi-vg/sqlgen
```

### Usage

```
Usage of sqlgen:
  -file string
    	input file name; required
  -o string
    	output file name
  -pkg string
    	output package name
```

### Tutorial

First, let's start with a simple `User` struct in `user.go`:

```Go
type User struct {
	ID     int64
	Login  string
	Email  string
}
```

We can run the following command:

```
sqlgen -file user.go -pkg demo
```

The tool outputs the following generated code:

```Go
func ScanUser(row *sql.Row) (*User, error) {
	var v0 int64
	var v1 string
	var v2 string

	err := row.Scan(
		&v0,
		&v1,
		&v2,
	)
	if err != nil {
		return nil, err
	}

	v := &User{}
	v.ID = v0
	v.Login = v1
	v.Email = v2

	return v, nil
}

// more functions and sql statements not displayed
```

### Nesting

Nested Go structures can be flattened into a single database table. As an example, we have a `User` and `Address` with a one-to-one relationship. In some cases, we may prefer to de-normalize our data and store in a single table, avoiding un-necessary joins.

```diff
type User struct {
    ID     int64  `sql:"pk: true"`
    Login  string
    Email  string
+   Addr   *Address
}

type Address struct {
    City   string
    State  string
    Zip    string `sql:"index: user_zip"`
}
```

The above relationship is flattened into a single table (see below). When the data is retrieved from the database the nested structure is restored.

```sql
CREATE TALBE IF NOT EXISTS users (
 user_id         INTEGER PRIMARY KEY AUTO_INCREMENT
,user_login      TEXT
,user_email      TEXT
,user_addr_city  TEXT
,user_addr_state TEXT
,user_addr_zip   TEXT
);
```

### JSON Encoding

Some types in your struct may not have native equivalents in your database such as `[]string`. These values can be marshaled and stored as JSON in the database.

```diff
type User struct {
    ID     int64  `sql:"pk: true"`
    Login  string
    Email  string
+   Label  []string `sql:"encode: json"
}
```

### Dialects

You may specify one of the following SQL dialects when generating your code: `postgres`, `mysql` and `sqlite`. The default value is `sqlite`.

```
sqlgen -file user.go -type User -pkg demo -db postgres
```


### Go Generate

Example use with `go:generate`:

```Go
package demo

//go:generate sqlgen -file user.go -type User -pkg demo -o user_sql.go

type User struct {
    ID     int64  `sql:"pk: true, auto: true"`
    Login  string `sql:"unique: user_login"`
    Email  string `sql:"size: 1024"`
    Avatar string
}
```

### Benchmarks

This tool demonstrates performance gains, albeit small, over light-weight ORM packages such as `sqlx` and `meddler`. Over time I plan to expand the benchmarks to include additional ORM packages.

To run the project benchmarks:

```
go get ./...
go generate ./...
go build
cd bench
go test -bench=Bench
```

Example selecing a single row:

```
BenchmarkMeddlerRow-4      30000        42773 ns/op
BenchmarkSqlxRow-4         30000        41554 ns/op
BenchmarkSqlgenRow-4       50000        39664 ns/op

```

Selecting multiple rows:

```
BenchmarkMeddlerRows-4      2000      1025218 ns/op
BenchmarkSqlxRows-4         2000       807213 ns/op
BenchmarkSqlgenRows-4       2000       700673 ns/op
```


#### Credits

This tool was inspired by [scaneo](https://github.com/variadico/scaneo).
