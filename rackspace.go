package main

import (
	"fmt"
	"log"

	"github.com/codegangsta/cli"
	"github.com/howeyc/gopass"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/rackspace"
)

//
// rackspaceAuth takes a cli context, user, pass, and key.
// Then it authenticates with Rackspace.
func rackspaceAuth(c *cli.Context, user, pass, key string) (*gophercloud.ProviderClient, error) {

	// We need authenication options, and an error.
	var ao gophercloud.AuthOptions
	var err error

	// If a user exists, but no password, then ask for it.
	if len(user) > 0 && len(pass) == 0 {

		// If the user wants verbosity give them verbosity...
		if c.Bool("verbose") {
			fmt.Println("A user was provided without a password at the command line...")
		}

		// Check if the user asked us to skip begging for a password.
		if c.GlobalBool("skip") {

			// If the user wants verbosity give them verbosity...
			if c.Bool("verbose") {
				fmt.Println("We would ask the user for a Rackspace password, but the skip flag was set to true...")
			}

		} else {

			// Tell the user we want the Rackspace Password.
			// This is helpful so that they do not think we are the proxy.
			fmt.Printf("Rackspace Password: ")
			pass = string(gopass.GetPasswdMasked()) // Masked passwords uses *s for characters.
		}
	}

	// No valid user.
	if len(user) == 0 {

		// If the user wants verbosity give them verbosity...
		if c.Bool("verbose") {
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
		if c.Bool("verbose") {
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
