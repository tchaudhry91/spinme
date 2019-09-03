# SpinMe!
[![Go Report Card](https://goreportcard.com/badge/github.com/tchaudhry91/spinme)](https://goreportcard.com/report/github.com/tchaudhry91/spinme)

A simple wrapper around Docker to quickly run "spin" up supporting containers such as databases for development.
SpinMe can be invoked via the CLI binary or used as a library straight from your Go code.

## CLI

Use this for daily use to start/stop/view your containers.

Usage: 
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

The password wherever needed is set to `password`. Once the container is up, you may modify it as required. Future versions will support supplying root passwords at create time (PRs welcome!)


## Library

This is a very early release, the eventual goal is to be able to use this system with CIs, to allow stuff like `go test` to create containers for databases on the fly and clean them up.

[GoDoc] (https://godoc.org/github.com/tchaudhry91/spinme/spin)


## Contributing

Contributions are very welcome. Please see create an issue!

## Licensing

The project is licensed under the Apache V2 License. See [License](LICENSE) for more information