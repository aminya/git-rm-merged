package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "git-rm-merged"
	app.Version = "1.0.0"
	app.Description = "The git sub command that removes the merged local branches."
	app.Action = mainAction

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func mainAction(c *cli.Context) error {
	branches, err := getMergedBranches()
	if err != nil {
		return err
	}

	for _, b := range branches {
		fmt.Printf("%s [y/n]:", b)

		var key string

		fmt.Scanln(&key)
		if (key != "y") {
			continue
		}

		err := delBranch(b)
		if err != nil {
			return err
		}
	}

	return nil
}
