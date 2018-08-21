package main

import (
	"errors"
	"os"

	"github.com/CityOfZion/neo-storm/compiler"
	"github.com/urfave/cli"
)

var (
	errNoInput = errors.New("No input file was found, specify an input file with the '--in or -i' flag")
)

func main() {
	ctl := cli.NewApp()
	ctl.Name = "neo-storm"
	ctl.Usage = "Neo smart contract framework for the Go programming language"
	ctl.Commands = []cli.Command{
		{
			Name:   "compile",
			Usage:  "compile a smart contract to an .avm file",
			Action: contactCompile,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "in, i",
					Usage: "input file to be compiled",
				},
				cli.StringFlag{
					Name:  "out, o",
					Usage: "output destination of the compiled contract",
				},
				cli.BoolFlag{
					Name:  "debug, d",
					Usage: "compile the contract in debug mode for additional compile information",
				},
			},
		},
	}
	ctl.Run(os.Args)
}

func contactCompile(ctx *cli.Context) error {
	src := ctx.String("in")
	if len(src) == 0 {
		return cli.NewExitError(errNoInput, 1)
	}

	o := &compiler.Options{
		Outfile: ctx.String("out"),
		Debug:   ctx.Bool("debug"),
	}

	if err := compiler.CompileAndSave(src, o); err != nil {
		return cli.NewExitError(err, 1)
	}

	return nil
}
