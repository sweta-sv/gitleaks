package main

import (
	_ "bufio"
	"fmt"
	"log"
	"os"
	"os/user"
	"strconv"
	_ "syscall"
)

const usage = `usage: gronit [arguments [options]] [options]
	
Arguments:
    start 		Starts gronit server
    restart 		Restarts gronit server
    stop 		Stops gronit server
	
Options:
    -p --port 		Port to run server on
    -h --help 		Display this message

`

const defaultPort = 3231

type System struct {
	CronPrefix string
	OS         string
	User       string
}

type Options struct {
	Start   bool
	Stop    bool
	Restart bool
	Port    int
}

// defaultSys fills a System struct with path to the crontab directory,
// default username, and type of system (macOS or Linux from `uname`)
func defaultSys() *System {
	var (
		cronPrefix string
		err        error
	)

	_user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	userName := _user.Username
	fmt.Println(userName)

	return &System{
		CronPrefix: cronPrefix,
		User:       userName,
	}
}

// help prints the usage string and exits
func help() {
	os.Stderr.WriteString(usage)
	os.Exit(1)
}

// optionsNextInt is a parseOptions helper that returns the value (int) of an option
// if valid.
func optionsNextInt(args []string, i *int) int {
	if len(args) > *i+1 {
		*i++
	} else {
		help()
	}
	argInt, err := strconv.Atoi(args[*i])
	if err != nil {
		fmt.Printf("Invalid %s option: %s\n", args[*i-1], args[*i])
		help()
	}
	return argInt
}

// optionsNextString is a parseOptions helper that returns the value (string) of an option
// if valid.
func optionsNextString(args []string, i *int) string {
	if len(args) > *i+1 {
		*i++
	} else {
		help()
	}
	return args[*i]
}

// parseOptions
func parseOptions(defaultSys *System, args []string) *Options {
	opts := &Options{}

	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "start":
			opts.Start = true
		case "restart":
			opts.Stop = true
		case "stop":
			opts.Restart = true
		case "-p", "--port":
			opts.Port = optionsNextInt(args, &i)
		case "-h", "--help":
			help()
			return nil
		default:
			fmt.Printf("Uknown option %s\n\n", arg)
			help()
			return nil
		}
	}

	if opts.Port == 0 {
		opts.Port = defaultPort
	}

	return opts
}
