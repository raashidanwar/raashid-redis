package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/urfave/cli/v2"
)

var config Config

func loadConfiguration() {
	configFile, err := os.Open("config.json")
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
}

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
						setUrl := config.REDIS_SEVER_URL + "set?key=" + key + "&value=" + value
						values := map[string]string{}
						json_data, err := json.Marshal(values)
						res, err := http.Post(setUrl, "application/json", bytes.NewBuffer(json_data))
						if err != nil {
							fmt.Printf("error making http request: %s\n", err)
							os.Exit(1)
						}
						resBody, err := ioutil.ReadAll(res.Body)
						fmt.Printf("%s\n", resBody)
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
						getUrl := config.REDIS_SEVER_URL + "get?key=" + key
						res, err := http.Get(getUrl)
						resBody, err := ioutil.ReadAll(res.Body)
						if err != nil {
							fmt.Printf("error making http request: %s\n", err)
							os.Exit(1)
						}
						fmt.Printf("%s\n", resBody)
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

func init() {
	loadConfiguration()
}

func main() {
	commands()
}
