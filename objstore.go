package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
)

// CURRENT USAGE
//
// objstore -rackspace upload -r -v path/to/directory (this command would look for env variables to authenticate)
// objstore -rackspace -user="username" -key="api key" upload -r -v path/to/directory
//

// FUTURE USAGE
//
// objstore [global options]      command [command options] [arguments...]
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
	app.Name = "objstore"
	app.Usage = "A CLI tool for uploading files to OpenStack object storage. Rackspace batteries included."
	app.Version = "0.0.1"
	app.Author = "Jack Spirou"
	app.Email = "jack.spirou@me.com"

	//
	// GLOBAL FLAGS
	//
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "rackspace, rack, rackspace_us, rack_us",
			Usage: "Rackspace Mode: Authenticate to US Rackspace.",
		},
		cli.BoolFlag{
			Name:  "rackspace_uk, rack_uk",
			Usage: "Rackspace Mode: Authenticate to UK Rackspace.",
		},
		cli.StringFlag{
			Name:  "u, user, username",
			Usage: "Authentication: The username to authenticate.",
		},
		cli.StringFlag{
			Name:  "p, pass, password",
			Usage: "Authentication: The password to authenticate.",
		},
		cli.StringFlag{
			Name:  "k, key, apikey, APIkey, APIKey",
			Usage: "Authentication: The API key to authenticate.",
		},
		cli.BoolFlag{
			Name:  "skip",
			Usage: "Authentication: It is assumed by default that we should ask you for your password if a user has been given, but a password is missing.  To skip this functionality, set skip to true.",
		},
	}

	//
	// COMMANDS
	//
	app.Commands = []cli.Command{

		//
		// Authenticate Command
		{
			Name:        "authenticate",
			Usage:       "Authenticate to your OpenStack instance or Rackspace account.",
			ShortName:   "auth",
			Description: "Use this command to test your credentials.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "v, verbose",
					Usage: "Verbose Mode: Show stuff as it happens.",
				},
			},
			Action: func(c *cli.Context) {
				_, err := auth(c)
				if err != nil {
					log.Fatalln("Unable to authenticate: " + err.Error())
				}
				fmt.Println("Authenticated Successfully!")
			},
		},

		//
		// Upload Command
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
