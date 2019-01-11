package main

import (
    "os"
    "github.com/urfave/cli"
)

func main() {
  app := cli.NewApp()

  app.Name = "freee_hr_cli"
  app.Usage = "This app is for freee hr"
  app.Version = "0.0.1"

  app.Run(os.Args)
}
