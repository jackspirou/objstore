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
// 						-rackspace            upload    -r              /path/to/file.txt
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

	//
	// CLI APP
	//
	app := cli.NewApp()

	//
	// CLI APP INFO
	//
	app.Name = "objstorage"
	app.Usage = "A CLI tool for uploading files to OpenStack object storage. Rackspace batteries included."
	app.Version = "0.0.1"
	app.Author = "Jack Spirou"
	app.Email = "jack.spirou@me.com"

	//
	// GLOBAL FLAGS
	//
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "rack, rackspace",
			Usage: "Rackspace Mode: Authenticate to US Rackspace.",
		},
		cli.BoolFlag{
			Name:  "rack_us, rackspace_us",
			Usage: "Rackspace Mode: Authenticate to US Rackspace.",
		},
		cli.BoolFlag{
			Name:  "rack_uk, rackspace_uk",
			Usage: "Rackspace Mode: Authenticate to UK Rackspace.",
		},
		cli.StringFlag{
			Name:  "u, user",
			Usage: "Authentication: The username to authenticate.",
		},
		cli.StringFlag{
			Name:  "p, pass",
			Usage: "Authentication: The password to authenticate.",
		},
		cli.StringFlag{
			Name:  "k, key",
			Usage: "Authentication: The API key to authenticate.",
		},
	}

	//
	// COMMANDS
	//
	app.Commands = []cli.Command{

		//
		// Authenticate command
		{
			Name:        "authenticate",
			ShortName:   "auth",
			Usage:       "Authenticate to your OpenStack instance or Rackspace account.",
			Description: "Use this command to test your credentials.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "r",
					Usage: "Recursive Mode: Recursively traverse a directory.",
				},
				cli.BoolFlag{
					Name:  "v, verbose",
					Usage: "Verboxe Mode: Show stuff as it happens.",
				},
			},
			Action: func(c *cli.Context) {
				println("Upload Stuff!")
			},
		},

		//
		// Upload command
		{
			Name: "upload",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "r",
					Usage: "Recursive Mode: Recursively traverse a directory.",
				},
				cli.BoolFlag{
					Name:  "v, verbose",
					Usage: "Verboxe Mode: Show stuff as it happens.",
				},
			},
			Action: func(c *cli.Context) {
				println("Upload Stuff!")
			},
		},
	}

	//
	// DEFAULT ACTION
	//
	app.Action = func(c *cli.Context) {
		cli.ShowAppHelp(c)
	}

	//
	// RUN APP
	//
	app.Run(os.Args)
}
