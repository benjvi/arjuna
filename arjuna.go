package main

import (
	"github.com/urfave/cli"
	"github.com/benjvi/arjuna/command"
	"os"
)

const usage =  `Intelligent inventory auditing and fixing for cloud resources

`

func main() {
	app := cli.NewApp()
	app.Name = "arjuna"
	app.Flags = []cli.Flag{}
	app.Commands = []cli.Commands{
		/*{
			Name:      "enforce",
			ShortName: "exec",
			Usage:     "Initialize a new project, creating a glide.yaml file",
			Description: `Runs specified policies in the current cloud environment (ie an account)

			Performs all functionality in sequence
			Fetch all resources of the types specified.
			Filters are applied to get unwanted or non-compliant resources.
			Individual actions can be taken on the resources to destroy or disable them
			Make assertions against the set of non-compliant resources
			Alert if assertion failed with details of actions applied`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "skip-import",
					Usage: "When initializing skip importing from other package managers.",
				},
				cli.BoolFlag{
					Name:  "non-interactive",
					Usage: "Disable interactive prompts.",
				},
			},
			Action: func(c *cli.Context) error {
				action.Create(".", c.Bool("skip-import"), c.Bool("non-interactive"))
				return nil
			},
		},*/
		{
			Name:      "audit",
			ShortName: "ls",
			Usage:     "",
			Description: `Gets all resources of types specified in current cloud environment

			Fetch all resources of the types specified.
			Filters are applied to get unwanted or non-compliant resources.
			Print out list of resources remaining`,
			Flags: []cli.Flag{
				cli.StringSliceFlag{
					Name:  "policies",
					Usage: "Limit policy files executed to those explicitly specified",
				},
			},
			Action: func(c *cli.Context) error {
				pwd, err := os.Getwd()
				if err != nil {
					return err
				}
				command.Audit(pwd)
				return nil
			},
		},/*
		{
			Name:      "alert",
			ShortName: "alert",
			Usage:     "",
			Description: `Gets all resources and alert if assertion fails

			Fetch all resources of the types specified.
			Filters are applied to get unwanted or non-compliant resources.
			Individual actions can be taken on the resources to destroy or disable them
			Make assertions against the set of non-compliant resources
			Alert if assertion failed with details of actions applied`,
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "skip-import",
					Usage: "When initializing skip importing from other package managers.",
				},
				cli.BoolFlag{
					Name:  "non-interactive",
					Usage: "Disable interactive prompts.",
				},
			},
			Action: func(c *cli.Context) error {
				action.Create(".", c.Bool("skip-import"), c.Bool("non-interactive"))
				return nil
			},
		},*/
	}
}




