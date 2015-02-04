# objstore
A CLI tool for uploading files to OpenStack object storage. Rackspace batteries included.

> Warning - This project is currently in initial development, expect problems.

## Installation Options

1. build this repository from source
2. `go get github.com/jackspirou/objstore`

> All options above will require you to have the compiled binary in your `$PATH`.

## Authentication

The good news is that the `objstore` binary makes many different ways of authenticating available to you.  The bad news is that you still have to authenticate.  The worse news is that ssh keys have no power here, so its all usernames, passwords, and API keys. 

### Rackspace

> Right now only Rackspace is supported.  OpenStack should be added later.

To authenticate with Rackspace, you can either enter your information at the command-line, or as enviorment variables.

