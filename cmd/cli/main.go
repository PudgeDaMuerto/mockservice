package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"os"

	"github.com/pudgedamuerto/mockservice/internal/app"
	"github.com/pudgedamuerto/mockservice/internal/config"
	"github.com/urfave/cli/v3"
)

var (
	version = "dev"
	commit  = "none"
)

func main() {
	listenFlag := &cli.StringFlag{
		Name:    "listen",
		Usage:   "address to listen on",
		Value:   "0.0.0.0:4000",
		Aliases: []string{"l"},
		Validator: func(addr string) error {
			_, _, err := net.SplitHostPort(addr)
			return err
		},
	}

	stdinFlag := &cli.BoolFlag{
		Name:  "stdin",
		Usage: "use config from stdin",
	}

	configArg := &cli.StringArg{
		Name: "config",
	}

	versionFlag := &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Action: func(ctx context.Context, c *cli.Command, b bool) error {
			fmt.Printf("version: %s\ncommit: %s\n", version, commit)
			os.Exit(0)
			return nil
		},
	}

	serveCmd := &cli.Command{
		Name:      "serve",
		Usage:     "start mockservice",
		Action:    serve,
		Flags:     []cli.Flag{listenFlag, stdinFlag},
		Arguments: []cli.Argument{configArg},
	}

	cmd := &cli.Command{
		Name:     "mockservice",
		Commands: []*cli.Command{serveCmd},
		Flags:    []cli.Flag{versionFlag},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "failed to serve:\n\t%v\n", err)
		os.Exit(1)
	}
}

func serve(ctx context.Context, cmd *cli.Command) error {
	isStdin := cmd.Bool("stdin")
	addr := cmd.String("listen")
	configPath := cmd.StringArg("config")

	if configPath == "" && !isStdin {
		return errors.New("config file is required. put it as second argument or from stdin using --stdin")
	}

	var file io.Reader = os.Stdin
	if !isStdin {
		var err error
		file, err = os.Open(configPath)
		if err != nil {
			return err
		}
	}

	serviceConfig, err := config.Load(file)
	if err != nil {
		return err
	}

	app := app.NewApp(*serviceConfig, addr)
	err = app.Serve()

	return err
}
