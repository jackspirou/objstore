package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codegangsta/cli"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
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
			Name:  "rackspace, rack, rackspace_us, rack_us",
			Usage: "Rackspace Mode: Authenticate to US Rackspace.",
		},
		cli.BoolFlag{
			Name:  "rackspace_uk, rack_uk",
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
		// Authenticate Command
		{
			Name:        "authenticate",
			Usage:       "Authenticate to your OpenStack instance or Rackspace account.",
			ShortName:   "auth",
			Description: "Use this command to test your credentials.",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "r, recursive",
					Usage: "Recursive Mode: Upload files by recursively traversing a directory.",
				},
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

//
// auth take a cli context and authenticates a user to OpenStack or Rackspace.
func auth(c *cli.Context) (*gophercloud.ProviderClient, error) {

	// Get authentication flags.
	user := c.GlobalString("user")
	pass := c.GlobalString("pass")
	key := c.GlobalString("key")

	// It is only possible to connect to Rackspace US or UK, not both.
	if c.GlobalBool("rackspace_us") && c.GlobalBool("rackspace_uk") {
		log.Fatalln("It is only possible to connect to Rackspace US or UK, not both.")
	}

	// Authenticate to Rackspace.
	if c.GlobalBool("rackspace_us") || c.GlobalBool("rackspace_uk") {

		// Authenticate with Rackspace.
		return authWithRackspace(c, user, pass, key)
	}
}

//
// authWithRackspace takes a cli context, user, pass, and key.
// Then it authenticates with Rackspace.
func authWithRackspace(c *cli.Context, user, pass, key string) (*gophercloud.ProviderClient, error) {

	// We need authenication options, and an error.
	var ao gophercloud.AuthOptions
	var err error

	// No valid user.
	if len(user) == 0 {

		// If the user wants verbosity give them verbosity...
		if c.GlobalBool("verbose") {
			fmt.Println("A user was not provided at the command line...")
			fmt.Println("Searching enviorment variables for authentication options...")
		}

		// Search enviorment variables.
		ao, err = rackspace.AuthOptionsFromEnv()

		// Handle the error, as go programmers should...
		if err != nil {
			log.Fatalln(err.Error())
		}

	} else if len(pass) == 0 && len(key) == 0 { // No password or API key.

		// If the user wants verbosity give them verbosity...
		if c.GlobalBool("verbose") {
			fmt.Println("A password or API key was not provided at the command line...")
			fmt.Println("Searching enviorment variables for authentication options...")
		}

		// Search enviorment variables.
		ao, err = rackspace.AuthOptionsFromEnv()

		// Handle the error, as go programmers should...
		if err != nil {
			log.Fatalln(err.Error())
		}

	} else { // We have enough information from the command line to authenticate.

		// Set available authentication options.
		ao.Username = user
		ao.Password = pass
		ao.APIKey = key
	}

	// Set the Rackspace US URL endpoint.
	if c.GlobalBool("rackspace_us") {
		ao.IdentityEndpoint = rackspace.RackspaceUSIdentity
	}

	// Set the Rackspace UK URL endpoint.
	if c.GlobalBool("rackspace_uk") {
		ao.IdentityEndpoint = rackspace.RackspaceUKIdentity
	}

	return rackspace.AuthenticatedClient(ao)
}
