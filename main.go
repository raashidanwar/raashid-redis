package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/urfave/cli/v2"
)

var store = map[string]string{}
var wg sync.WaitGroup

func commands() {
	for len(os.Args) != 0 {
		app := &cli.App{
			Commands: []*cli.Command{
				{
					Name:    "set",
					Aliases: []string{"s"},
					Usage:   "set key, value pair",
					Action: func(cCtx *cli.Context) error {
						if cCtx.Args().Len() != 2 {
							return fmt.Errorf("expected 2 arguments {key, value}")
						}
						key := cCtx.Args().Get(0)
						value := cCtx.Args().Get(1)
						store[key] = value
						fmt.Printf("set %s to %s\n", key, value)
						return nil
					},
				},
				{
					Name:    "get",
					Aliases: []string{"g"},
					Usage:   "get key, value pair",
					Action: func(cCtx *cli.Context) error {
						if cCtx.Args().Len() != 1 {
							return fmt.Errorf("expected 1 argument {key}")
						}
						key := cCtx.Args().Get(0)
						value, err := store[key]
						if !err {
							return fmt.Errorf("key %s not found", key)
						}
						fmt.Printf("%s\n", value)
						return nil
					},
				},
			},
		}

		if err := app.Run(os.Args); err != nil {
			log.Fatal(err)
		}
		os.Args = []string{}
	}
}

func main() {
	commands()
}
