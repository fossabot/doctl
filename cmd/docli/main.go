package main

import (
	"os"

	"code.google.com/p/goauth2/oauth"

	"github.com/codegangsta/cli"
	"github.com/digitalocean/godo"
)

func main() {
	app := cli.NewApp()
	app.Name = "docli"
	app.Usage = "DigitalOcean API CLI"
	app.Version = "0.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "token",
			Usage:  "DigitalOcean API V2 Token",
			EnvVar: "DO_TOKEN",
		},
	}

	app.Commands = []cli.Command{
		dropletCommands(),
	}

	app.Run(os.Args)
}

func tokenFlag() cli.Flag {
	return cli.StringFlag{
		Name:   "token",
		Usage:  "DigitalOcean API V2 Token",
		EnvVar: "DO_TOKEN",
	}
}

func newClient(token string) *godo.Client {
	t := &oauth.Transport{
		Token: &oauth.Token{AccessToken: token},
	}

	return godo.NewClient(t.Client())
}