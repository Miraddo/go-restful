package main

import (
	"github.com/urfave/cli"
	"os"
	
	"log"
)

func main() {
	app := cli.NewApp()

	// add flags with three arguments
	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "save",
			Value: "no",
			Usage: "Should save to database (yes/no)",
		},
	}

	app.Version = "1.0"
	// This function parses and brings data in cli.Context struct
	app.Action = func(c *cli.Context) error {
		var args []string

		if c.NArg() > 0 {
			// Fetch arguments in a array
			args = c.Args()
			personName := args[0]
			marks := args[1:len(args)]
			log.Println("Person: ", personName)
			log.Println("marks", marks)
		}
		// check the flag value
		if c.String("save") == "no" {
			log.Println("Skipping saving to the database")
		} else {
			// Add database logic here
			log.Println("Saving to the database", args)
		}
		return nil
	}
	// Pass os.Args to cli app to parse content
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		return
	}
}