package main

import (
	"context"
	"fmt"
	"os"

	"code"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		Name:      "hexlet-path-size",
		Usage:     "print size of a file or directory",
		ArgsUsage: "<path>",
		Flags: []cli.Flag{
			&cli.BoolFlag{Name: "human",
				Aliases: []string{"H"},
				Usage:   "human-readable sizes (auto-select unit)",
			},

			&cli.BoolFlag{
				Name:    "all",
				Aliases: []string{"a"},
				Usage:   "include hidden files and directories",
			},

			&cli.BoolFlag{
				Name:    "recursive",
				Aliases: []string{"r"},
				Usage:   "recursive size of directories",
			},
		},

		Action: func(ctx context.Context, cmd *cli.Command) error {
			path := cmd.Args().First()
			human := cmd.Bool("human")
			all := cmd.Bool("all")
			recursive := cmd.Bool("recursive")

			size, err := code.GetPathSize(path, recursive, human, all)
			if err != nil {
				return err
			}

			fmt.Printf("%s\t%s\n", size, path)

			return nil
		},
	}

	err := cmd.Run(context.Background(), os.Args)
	if err != nil {
		panic(err)
	}
}
