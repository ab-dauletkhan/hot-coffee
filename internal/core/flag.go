package core

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var (
	Port int
	Dir  string
	Help bool
	Env  string
)

func ParseFlags() error {
	flag.IntVar(&Port, "port", 8080, "port to bind the server to")
	flag.StringVar(&Dir, "dir", "./data", "directory to serve static files from")
	flag.BoolVar(&Help, "help", false, "display help message")
	flag.StringVar(&Env, "env", "local", "environment to run the server in, accepted values are: 'local', 'dev', 'prod'")

	flag.Usage = printUsage
	flag.Parse()

	if Help {
		printUsage()
		os.Exit(0)
	}

	if Env != "local" && Env != "dev" && Env != "prod" {
		return fmt.Errorf("invalid environment: %s, accepted values are: 'local', 'dev', 'prod'", Env)
	}

	filepath.Clean(Dir)
	if Port < 1024 || Port > 49151 {
		return fmt.Errorf("invalid port number: %d, accepted range is 1024 - 49151", Port)
	}
	return nil
}

func printUsage() {
	fmt.Println(`$ ./hot-coffee --help
Coffee Shop Management System

Usage:
  hot-coffee [--port <N>] [--dir <S>]
  hot-coffee --help

Options:
  --help       Show this screen.
  --port N     Port number.
  --dir S      Path to the data directory.`)
}
