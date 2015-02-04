package main

import (
	"log"

	"github.com/codegangsta/cli"
	"github.com/rackspace/gophercloud"
)

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
		return rackspaceAuth(c, user, pass, key)
	}

	// We only support Rackspace right now... sorry... :(
	log.Fatalln("We only support Rackspace authentication right now.")
	return nil, nil
}
