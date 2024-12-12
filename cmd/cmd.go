package cmd

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	cli "github.com/urfave/cli/v2"
)

func Execute(args []string) {
	ctx, cancel := context.WithCancel(context.Background())

	go runSyscallContextCancellation(ctx, cancel)

	cliApp := &cli.App{
		Name:  "weather-cli",
		Usage: "weather-cli",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "debug",
				Aliases:     []string{"d"},
				Usage:       "enable debug mode, which will print data from weather clients",
				DefaultText: "false",
			},
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.yaml",
			},
		},
	}

	cliApp.Commands = []*cli.Command{
		run(),
	}

	err := cliApp.RunContext(ctx, args)
	if err != nil {
		cancel()
		log.Fatalf("execute failed: %v\n", err)
	}

	time.Sleep(time.Second)
}

func runSyscallContextCancellation(ctx context.Context, cancel context.CancelFunc) {
	signalCh := make(chan os.Signal, 1)
	signal.Notify(
		signalCh,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)

	func() {
		select {
		case <-signalCh:
			cancel()
			return
		case <-ctx.Done():
			return
		}
	}()
}
