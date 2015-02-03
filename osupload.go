package main

import (
	"os"

	"github.com/codegangsta/cli"
)

// USAGE
//
// objstorage [global options]      command [command options] [arguments...]
//            -rackspace            upload    -r              /path/to/file.txt
//            -rackspace_us         update    -v              /path/to/directory
//            -rackspace_uk         delete
//            -user="username"      list
//            -pass="password"
//            -key="api key"
//            -container="container name"
//            -account="account number" or tenantID="tenant id"
//            -identityendpoint="identity endpoint"
//

// objstorage -rackspace -r -v -user -pass -key -container /path/to/file.txt
// objstorage -rackspace -user -pass -key -container /path/to/file.txt
// objstorage -user -pass -key -container /path/to/file.txt
//

func main() {

	// Create a new CLI app.
	app := cli.NewApp()

	// Set CLI Info.
	app.Name = "osupload"
	app.Usage = "A CLI tool for uploading files to OpenStack object storage."
	app.Version = "0.0.1"
	app.Author = "Jack Spirou"
	app.Email = "jack.spirou@me.com"

	// Set CLI Flags.
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "config, c",
			Value: "config.json",
			Usage: "The path to the config file. Defaults to 'config.json' in the execution path.",
		},
		cli.StringFlag{
			Name:  "env, e",
			Value: "dev",
			Usage: "The name of the environment to be used",
		},
	}

	// Its action time.
	app.Action = func(c *cli.Context) {
		println("upload! I say!")
	}

	// Run the CLI app with whatever args may come.
	app.Run(os.Args)
}
