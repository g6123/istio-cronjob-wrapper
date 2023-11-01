package main

import (
	"os"

	"github.com/g6123/istio-cronjob-wrapper/pkg"
	"github.com/op/go-logging"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name: pkg.Name,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "verbose",
				Aliases: []string{"v"},
			},
			&cli.BoolFlag{
				Name:    "force",
				Aliases: []string{"f"},
			},
			&cli.StringFlag{
				Name:  "ready",
				Value: "http://localhost:15020/healthz/ready",
			},
			&cli.IntFlag{
				Name:  "connect-timeout",
				Value: 1,
			},
			&cli.IntFlag{
				Name:  "max-retry",
				Value: 300,
			},
			&cli.StringFlag{
				Name:  "quit",
				Value: "http://localhost:15000/quitquitquit",
			},
		},
		Action: func(ctx *cli.Context) error {
			if ctx.NArg() == 0 {
				return cli.Exit("no argumets", 1)
			}

			if ctx.Bool("verbose") {
				logging.SetLevel(logging.DEBUG, pkg.Name)
			} else {
				logging.SetLevel(logging.WARNING, pkg.Name)
			}

			if !ctx.Bool("force") && !pkg.IsKube() {
				return pkg.Exec(ctx.Args())
			}

			err := pkg.WaitEnvoyReady(ctx.String("ready"), ctx.Int("connect-timeout"), ctx.Int("max-retry"))
			if err != nil {
				return err
			}
			pkg.Logger.Infof("envoy ready")

			code, err := pkg.Run(ctx.Args())
			if err != nil {
				return err
			}
			if code < 0 {
				pkg.Logger.Warningf("process terminated by external signal")
			}
			if code > 0 {
				pkg.Logger.Warningf("process exited with non-zero exit code %d", code)
			}

			err = pkg.KillEnvoy(ctx.String("quit"))
			if err != nil {
				return err
			}
			pkg.Logger.Infof("envoy quit")

			return cli.Exit("", code)
		},
	}

	if err := app.Run(os.Args); err != nil && err.Error() != "" {
		pkg.Logger.Fatal(err)
	}
}
