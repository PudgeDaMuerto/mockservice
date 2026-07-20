package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/pudgedamuerto/mockservice/internal/app"
	"github.com/pudgedamuerto/mockservice/internal/config"
	"github.com/urfave/cli/v3"
)

func main() {
	listenFlag := &cli.StringFlag{
		Name:    "listen",
		Usage:   "address to listen on",
		Value:   ":3000",
		Aliases: []string{"l"},
		Validator: func(addr string) error {
			_, _, err := net.SplitHostPort(addr)
			return err
		},
	}

	configArg := &cli.StringArg{
		Name: "config",
	}

	serveCmd := &cli.Command{
		Name:      "serve",
		Usage:     "start mockservice",
		Action:    serve,
		Flags:     []cli.Flag{listenFlag},
		Arguments: []cli.Argument{configArg},
	}

	cmd := &cli.Command{
		Name:     "mockservice",
		Commands: []*cli.Command{serveCmd},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "failed to serve:\n\t%v\n", err)
		os.Exit(1)
	}
}

func serve(ctx context.Context, cmd *cli.Command) error {
	configPath := cmd.StringArg("config")
	if configPath == "" {
		return errors.New("config file is required")
	}

	addr := cmd.String("listen")

	serviceConfig, err := config.Load(configPath)
	if err != nil {
		return err
	}

	app := app.NewApp(*serviceConfig, addr)

	err = app.Serve()

	return err
}
