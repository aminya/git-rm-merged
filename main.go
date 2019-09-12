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
  app.Description = "The git sub command that removes the merged local branches."
  app.Action = func(c *cli.Context) error {
    fmt.Println("git-rm-merged!")
    return nil
  }

  err := app.Run(os.Args)
  if err != nil {
    log.Fatal(err)
  }
}
