package main

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/CityOfZion/neo-go/pkg/rpc"
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
			Action: contractCompile,
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
		{
			Name:   "testinvoke",
			Usage:  "testinvoke a smart contract against a remote NEO RPC node",
			Action: contractTestInvoke,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "in, i",
					Usage: "input file to be compiled",
				},
			},
		},
	}
	ctl.Run(os.Args)
}

func contractCompile(ctx *cli.Context) error {
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

func contractTestInvoke(ctx *cli.Context) error {
	src := ctx.String("in")
	if len(src) == 0 {
		return cli.NewExitError(errNoInput, 1)
	}

	b, err := ioutil.ReadFile(src)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	endpoint := "http://seed2.ngd.network:10332"
	opts := rpc.ClientOptions{}
	client, err := rpc.NewClient(context.TODO(), endpoint, opts)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	scriptHex := hex.EncodeToString(b)
	resp, err := client.InvokeScript(scriptHex)
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	b, err = json.MarshalIndent(resp.Result, "", "  ")
	if err != nil {
		return cli.NewExitError(err, 1)
	}

	fmt.Println(string(b))

	return nil
}
