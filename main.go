package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sync"
)

func main() {
	app := &cli.App{
		Name:  "HashHelper",
		Usage: "A tool to help with the hash",
		Commands: []*cli.Command{
			{
				Name:    "Generate",
				Aliases: []string{"g"},
				Usage:   "Generate [filename].fs256",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "path",
						Aliases:     []string{"p"},
						Usage:       "specify path",
						Value:       "./",
						DefaultText: "./",
					},
				},
				Action: func(c *cli.Context) error {
					oripath := c.String("path")
					var wg sync.WaitGroup
					Generate(oripath, &wg)
					wg.Wait()
					return nil
				},
			},
			{
				Name:    "Verify",
				Aliases: []string{"v"},
				Usage:   "Verify [filename].fs256",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "path",
						Aliases:     []string{"p"},
						Usage:       "specify path",
						Value:       "./",
						DefaultText: "./",
					},
					&cli.StringFlag{
						Name:        "remove",
						Aliases:     []string{"r"},
						Usage:       "remove .fs256 file after verify",
						Value:       "false",
						DefaultText: "false",
					},
				},
				Action: func(c *cli.Context) error {
					oripath := c.String("path")
					temp := c.Bool("remove")
					isError := false
					var wg sync.WaitGroup
					Verify(oripath, &wg, &isError)
					wg.Wait()
					if temp && !isError {
						err := DeleteHashFiles(oripath)
						if err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
