package pkg

import (
	"errors"
	"os"
	"os/exec"
	"syscall"

	"github.com/urfave/cli/v2"
)

func Run(args cli.Args) (int, error) {
	cmd := exec.Command(args.First(), args.Tail()...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		if exit_err := (&exec.ExitError{}); errors.As(err, &exit_err) {
			return exit_err.ExitCode(), nil
		} else {
			return 0, err
		}
	}

	return 0, nil
}

func Exec(args cli.Args) error {
	argv0, err := exec.LookPath(args.First())
	if err != nil {
		return err
	}

	err = syscall.Exec(argv0, args.Slice(), os.Environ())
	if err != nil {
		return err
	}

	return nil
}
