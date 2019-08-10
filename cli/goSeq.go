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
	app.Name = "GoSeq"
	app.Usage = "CLI for calling GoSeq Sequence Analysis Algorithmn"
	app.Author = "Lukas Valentin Micke"
	app.Version = "0.0.2"
}

func commands() {
	var alphabet string
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
		{
			Name:      "qGramDistance",
			Aliases:   []string{"qd", "qgramd"},
			Usage:     "Compute the q-gram Distance between two strings",
			ArgsUsage: "[args]",
			Flags: []cli.Flag{
				cli.BoolFlag{Name: "forever, forevvarr"},
				cli.StringFlag{Name: "alphabet, a", Usage: "Which alphabet to use: choice between: dna, rna, prot", Required: true,
					Value: "dna", Destination: &alphabet},
			},
			Action: func(c *cli.Context) {
				seqA := c.Args()[0]
				seqB := c.Args()[1]
				fmt.Println(distances.QGramDistance(seqA, seqB, alphabet))
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
