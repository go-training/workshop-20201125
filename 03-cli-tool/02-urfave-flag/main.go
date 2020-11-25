package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type Server struct {
	Language string
	Debug    bool
}

func main() {
	cfg := &Server{}
	app := &cli.App{
		Name:  "app",
		Usage: "fight the loneliness!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "lang",
				Value:       "english",
				Usage:       "language for the greeting",
				Destination: &cfg.Language,
			},
			&cli.BoolFlag{
				Name:        "debug",
				Value:       false,
				Usage:       "debug mode",
				Destination: &cfg.Debug,
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			if c.String("lang") == "spanish" {
				fmt.Println("Hola foobar")
			} else {
				fmt.Println("Hello foobar")
			}
			fmt.Println(cfg.Language)
			fmt.Println(cfg.Debug)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
