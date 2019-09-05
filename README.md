# SpinMe!
[![Go Report Card](https://goreportcard.com/badge/github.com/tchaudhry91/spinme)](https://goreportcard.com/report/github.com/tchaudhry91/spinme)
[![CI](https://github.com/tchaudhry91/spinme/workflows/Continuous%20Integration%20Workflow/badge.svg)](https://github.com/tchaudhry91/spinme)

A simple wrapper around Docker to quickly run "spin" up supporting containers such as databases for development.
SpinMe can be invoked via the CLI binary or used as a library straight from your Go code.

## Library

The primary goal of this library is to eliminate the need to externally spin up testing databases (esp while running tests on a CI system). 
You can create live docker containers straight from your go test files, making `go test` independent. This allows you to test against real databases, without having to mock anything.

See examples for Postgres/MySQL/Mongo/Redis in the [GoDoc](https://godoc.org/github.com/tchaudhry91/spinme/spin)

Sample Usage:
```
package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tchaudhry91/spinme/spin"
)

func main() {
	out, err := spin.MySQL(context.Background(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer spin.SlashID(context.Background(), out.ID)
	// Give mysql a minute to boot-up, sadly there is no "ready" check yet
	time.Sleep(1 * time.Minute)
	connStr, err := spin.MySQLConnString(out)
	if err != nil {
		fmt.Println(err)
		return
	}
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
}
```

## CLI

Use this for daily use to start/stop/view your containers.

### CLI Usage: 
```
spinme -h
SpinMe is a wrapper around docker to run common applications.
  Use this to easily create dependent services such as databases.

Usage:
  spinme [command]

Available Commands:
  down        Bring down the given container
  help        Help about any command
  status      Status shows the list of all running services spun via spinme
  up          Start a particular service

Flags:
      --db string   Database for local storage (default "/home/tchaudhry/.spinme")
  -h, --help        help for spinme

Use "spinme [command] --help" for more information about a command.
```

e.g

The tool currently contains the following dbs with some standard defaults that can be started with a single command such as
- `spinme up -s postgres`
- `spinme up -s mysql`
- `spinme up -s redis`
- `spinme up -s mongo`

Admin users are as follows:
- Postgres = `postgres`
- Mysql = `root`
- Mongo = `mongoadmin`

The password wherever needed is set to `password`. Once the container is up, you may modify it as required. The environment variables for the images to override this can also be set via the command line:
e.g `--env PG_PASSWORD=1231`


## Contributing

Contributions are very welcome. Please see create an issue!

## Licensing

The project is licensed under the Apache V2 License. See [License](LICENSE) for more information