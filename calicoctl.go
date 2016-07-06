package main

import (
	"flag"
	"fmt"

	"github.com/docopt/docopt-go"
	"github.com/projectcalico/libcalico/commands"
	"os"
)

func main() {
	usage := `Usage: calicoctl <command> [<args>...]

    create         Create a resource by filename or stdin.
    replace        Replace a resource by filename or stdin.
    delete         Delete a resource identified by file, stdin or resource type and name.
    get            Get a resource identified by file, stdin or resource type and name.
    version        Display the version of calicoctl.

See 'calicoctl <command> --help' to read about a specific subcommand.`
	var err error;
	doc := commands.EtcdIntro + usage

	if os.Getenv("GLOG") != "" {
		flag.Parse()
		flag.Lookup("logtostderr").Value.Set("true")
		flag.Lookup("v").Value.Set("10")
	}

	arguments, _ := docopt.Parse(doc, nil, true, "calicoctl", true, false)

	if arguments["<command>"] != nil {
		command := arguments["<command>"].(string)
		args := append([]string{command}, arguments["<args>"].([]string)...)

		switch command {
		case "create":
			err = commands.Create(args)
		case "replace":
			err = commands.Replace(args)
		case "delete":
			err = commands.Delete(args)
		case "get":
			err = commands.Get(args)
		case "version":
			err = commands.Version(args)
		default:
			fmt.Println(usage)
		}
	}

	if err != nil {
		os.Exit(1)
	}
}
