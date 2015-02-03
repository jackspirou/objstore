package main

import (
	"os"

	"github.com/codegangsta/cli"
)

// CURRENT USAGE
//
// objstorage -rackspace upload -r -v path/to/directory (this command would look for env variables to authenticate)
// objstorage -rackspace -user="username" -key="api key" upload -r -v path/to/directory
//

// FUTURE USAGE
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

func main() {

	// Create a new CLI app.
	app := cli.NewApp()

	// Set CLI Info.
	app.Name = "osupload"
	app.Usage = "A CLI tool for uploading files to OpenStack object storage. Rackspace batteries included."
	app.Version = "0.0.1"
	app.Author = "Jack Spirou"
	app.Email = "jack.spirou@me.com"

	// Define the CLI Flags.
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "rackspace",
			Usage: "Rackspace Mode: Authenticate to US Rackspace.",
		},
		cli.BoolFlag{
			Name:  "rackspace_us",
			Usage: "Rackspace Mode: Authenticate to US Rackspace.",
		},
		cli.BoolFlag{
			Name:  "rackspace_uk",
			Usage: "Rackspace Mode: Authenticate to UK Rackspace.",
		},
		cli.StringFlag{
			Name:  "user",
			Usage: "Authentication: The username used to authenticate.",
		},
		cli.StringFlag{
			Name:  "pass",
			Usage: "Authentication: The password used to authenticate.",
		},
		cli.StringFlag{
			Name:  "key",
			Usage: "The name of the environment to be used",
		},
	}

	// Define the CLI Commands.
	app.Commands = []cli.Command{
		{
			Name: "upload",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "r",
					Usage: "Recursive Mode: Recursively traverse a directory.",
				},
			},
		},
	}

	// Its action time.
	app.Action = func(c *cli.Context) {
		println("No command supplied.")
	}

	// Run the CLI app with whatever args may come.
	app.Run(os.Args)
}
