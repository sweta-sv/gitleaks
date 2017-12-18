package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"
	_ "syscall"
)

const usage = `usage: gronit [cmd] &/|| [options]

  Start: $ gronit start [options]
    -u --user	Select which user to update cron for
    -p --port	Which port to start service on

  Restart: $ gronit restart [options]

  Stop: $ gronit start [options]
  
  Add: $ gronit [options]
    -u --user		Select which user to update cron
    --loadyaml		Yaml file that contains a schedule
    --loadjson		JSON file that contains a schedule
    --loadcron		cron file that contains a schedule
    --path        	path to crontab:
	  * default cron path for osx: /var/at/tabs/$USER
	  * default cron path for linux: /etc/cron.d/$USER

  Display: $ gronit [options]
    --list-json 	Sends crontabs to stdout in json format
    --list-yaml 	Sends crontabs to stdout in yaml format
    -l, --list    	Sends crontabs to stdout in human readable form

  Interfacing: $ gronit [options]
    --remove=N[,..]	Remove gronit task(s)
    -v --version	Version

`

func init() {
	flag.Usage = func() {
		fmt.Printf(usage)
		flag.PrintDefaults()
	}
}

type System struct {
	CronPrefix string
	OS         string
	User       string
}

type Options struct {
	User     string
	Port     int
	LoadYaml string
	LoadJson string
	LoadCron string
}

type Task struct {
	DoEverySecond int
	DoEveryMinute int
	DoEveryHour   int
	DoEveryDay    int
	DoEveryMonth  int
	Cmd           string
	Heartbeat     bool
	Monitor       bool
}

func defaultSys() *System {
	var (
		cronPrefix string
		uname      []byte
		err        error
	)

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	usrName := usr.Username

	if uname, err = exec.Command("uname").Output(); err != nil {
		log.Fatal(err)
	}
	unameStr := strings.TrimSpace(string(uname))

	if unameStr == "Darwin" {
		cronPrefix = "/var/at/tabs/"

	} else if unameStr == "Linux" {
		cronPrefix = "/etc/cron.d/"
	}

	return &System{
		CronPrefix: cronPrefix,
		OS:         unameStr,
		User:       usrName,
	}
}

func help() {
	os.Stderr.WriteString(usage)
	os.Exit(1)
}

func parseOptions(args []string) *Options {
	if len(os.Args) < 2 {
		help()
	}
	opts := &Options{}

	// User     string
	// Port     int
	// LoadYaml string
	// LoadJson string
	// LoadCron string

	loadFile := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		switch arg {
		case "-u", "--user":
			i++
			opts.User = args[i]
		case "-p", "--port":
			i++
			port, err := strconv.Atoi(args[i])
			if err != nil {
				fmt.Printf("Invalid port option: %s\n", args[i])
				help()
				return nil
			}
			opts.Port = port
		case "--LoadYaml":
			i++
			if !(loadFile) {
				opts.LoadYaml = args[i]
				loadFile = true
			}
		case "--LoadJson":
			i++
			if !(loadFile) {
				opts.LoadJson = args[i]
				loadFile = true
			}
		case "--LoadCront":
			i++
			if !(loadFile) {
				opts.LoadCron = args[i]
				loadFile = true
			}
		default:
			fmt.Printf("Uknown option %s\n\n", arg)
			help()
			return nil
		}
	}

	fmt.Println(opts)

	return opts
}

func cronPrefix() string {
	return ""
}

func yamlToTask() *Task {
	return &Task{}
}

func jsonToTask() *Task {
	return &Task{}
}

func cronToTask() *Task {
	return &Task{}
}