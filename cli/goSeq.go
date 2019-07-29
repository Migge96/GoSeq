package main

import (
	"GoSeq/distances"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

var app = cli.NewApp()

func info() {
	app.Name = "GoSeq CLI"
	app.Usage = "CLI for calling GoSeq Sequence Analysis Algorithmn"
	app.Author = "Lukas Valentin Micke"
	app.Version = "0.0.1"
}

func commands() {
	app.Commands = []cli.Command{
		{
			Name:    "editDistance",
			Aliases: []string{"ed"},
			Usage:   "Compute the edit Distance betweeen two strings",
			Action: func(c *cli.Context) {
				seqA := c.Args()[0]
				seqB := c.Args()[1]
				fmt.Println(distances.EditDistance(seqA, seqB))
			},
		},
	}
}

func main() {
	info()
	commands()
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
